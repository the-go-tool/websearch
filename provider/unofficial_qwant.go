package provider

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"math"
	"net/url"
	"strconv"
	"sync"
	"websearch/helpers"
	"websearch/provider/errors"
)

// The Unofficial UnofficialQwant provider name
const ProviderUnofficialQwant = ProviderName("unofficial_qwant")

// The Unofficial UnofficialQwant [https://qwant.com] web search provider
type UnofficialQwant struct {
	api           url.URL
	defaultParams map[string]string
	locale        string
}

// The config for Unofficial Qwant provider
type UnofficialQwantConfig struct {
	Locale string
}

// Makes a new UnofficialQwant web search provider
func NewUnofficialQwant(config ...UnofficialQwantConfig) UnofficialQwant {
	conf := UnofficialQwantConfig{}
	if len(config) > 0 {
		conf = config[0]
	}

	if conf.Locale == "" {
		conf.Locale = "en_US"
	}

	api := url.URL{
		Scheme: "https",
		Host:   "api.qwant.com",
		Path:   "/api/search/web",
	}
	defaultParams := map[string]string{
		"t":   "web",
		"uiv": "4",
	}
	return UnofficialQwant{
		api:           api,
		defaultParams: defaultParams,
		locale:        conf.Locale,
	}
}

// Makes web search
func (engine UnofficialQwant) Search(query string, maxCount ...int) (Results, error) {
	const maxCountPerPage = 10

	count := maxCountPerPage
	if len(maxCount) > 0 {
		count = maxCount[0]
	}

	if count < 1 || count > maxCountPerPage * 5 {
		return nil, fmt.Errorf("incorrect count %d, expect value [1...50]", count)
	}

	// Results in common format as a parts to save ordering
	resultsMutex := sync.Mutex{}
	resultsParts := make([]Results, int(math.Ceil(float64(count)/10)))

	// Runs request for every page async
	group := errgroup.Group{}
	offset := 0
	links := count
	part := 0
	for links > 0 {
		func(offset, links, part int) {
			group.Go(func() error {
				// Request for page
				res, err := engine.search(query, offset, int(math.Min(maxCountPerPage, float64(links))))
				if err != nil {
					return err
				}

				// Converting to common results
				results := make(Results, 0, len(res.Data.Result.Items))
				for _, item := range res.Data.Result.Items {
					u, err := url.Parse(item.Url)
					if err != nil {
						return err
					}
					results = append(results, Result{
						Title:       item.Title,
						Description: item.Desc,
						Link:        *u,
					})
				}

				// Saving to result parts
				resultsMutex.Lock()
				resultsParts[part] = results
				resultsMutex.Unlock()

				return nil
			})
		}(offset, links, part)

		offset += maxCountPerPage
		links -= maxCountPerPage
		part++
	}

	// Waiting for requests and handling error
	if err := group.Wait(); err != nil {
		return nil, err
	}

	// Connecting results from all parts
	results := make(Results, 0, len(resultsParts)*maxCountPerPage)
	for i := range resultsParts {
		results = append(results, resultsParts[i]...)
	}

	// Checks for count
	if len(results) > count {
		results = results[:count-1]
	}

	return results, nil
}

// Returns provider name
func (engine UnofficialQwant) Name() ProviderName {
	return ProviderUnofficialQwant
}

// Inner searcher with pagination
func (engine UnofficialQwant) search(query string, offset int, count int) (unofficialQwantResults, error) {
	if count > 10 || count < 1 {
		return unofficialQwantResults{}, fmt.Errorf("incorrect count %d, expect value [1...10]", count)
	}
	if offset > 40 || offset < 0 {
		return unofficialQwantResults{}, fmt.Errorf("incorrect offset %d, expect value [0...40]", offset)
	}

	// Merges default params and external
	params := map[string]string{
		"q":      query,
		"count":  strconv.Itoa(count),
		"offset": strconv.Itoa(offset),
	}
	params = helpers.ParamsMerge(engine.defaultParams, params)

	// Makes full url
	u := engine.api
	u.RawQuery = helpers.ParamsRender(params)

	// Request
	results := unofficialQwantResults{}
	if err := helpers.RequestJSON(&results, u); err != nil {
		return unofficialQwantResults{}, err
	}

	// Handling UnofficialQwant errors
	switch results.Data.ErrorCode {
	case 0: //< All is ok
	case 429:
		return unofficialQwantResults{}, errors.NewIPBanned(fmt.Errorf("%d", results.Data.ErrorCode))
	case 14:
		return unofficialQwantResults{}, errors.NewBadRequestError(fmt.Errorf("%d", results.Data.ErrorCode))
	default:
		return unofficialQwantResults{}, fmt.Errorf("%d", results.Data.ErrorCode)
	}

	// Checks UnofficialQwant inner error
	if results.Data.ErrorCode != 0 {
		return unofficialQwantResults{}, fmt.Errorf("qwant error: %d", results.Data.ErrorCode)
	}

	return results, nil
}

// Results of UnofficialQwant search
type unofficialQwantResults struct {
	Data struct {
		Result struct {
			Items []struct {
				Desc  string `json:"desc"`
				Title string `json:"title"`
				Url   string `json:"url"`
			} `json:"items"`
		} `json:"result"`
		ErrorCode int `json:"error_code"`
	} `json:"data"`
}
