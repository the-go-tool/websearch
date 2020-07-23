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

// The Qwant provider name
const ProviderQwant = ProviderName("qwant")

// The Qwant [https://qwant.com] web search provider
type Qwant struct {
	api           url.URL
	defaultParams map[string]string
	locale        string
}

// Makes a new Qwant web search provider
func NewQwant(locale string) Qwant {
	api := url.URL{
		Scheme: "https",
		Host:   "api.qwant.com",
		Path:   "/api/search/web",
	}
	defaultParams := map[string]string{
		"t":   "web",
		"uiv": "4",
	}
	return Qwant{
		api:           api,
		defaultParams: defaultParams,
		locale:        locale,
	}
}

// Makes web search
func (engine Qwant) Search(query string, maxCount ...int) (Results, error) {
	const MAX_COUNT_PER_PAGE = 10

	count := MAX_COUNT_PER_PAGE
	if len(maxCount) > 0 {
		count = maxCount[0]
	}

	if count < 1 || count > 50 {
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
				res, err := engine.search(query, offset, int(math.Min(MAX_COUNT_PER_PAGE, float64(links))))
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

		offset += MAX_COUNT_PER_PAGE
		links -= MAX_COUNT_PER_PAGE
		part++
	}

	// Waiting for requests and handling error
	if err := group.Wait(); err != nil {
		return nil, err
	}

	// Connecting results from all parts
	results := make(Results, 0, len(resultsParts)*MAX_COUNT_PER_PAGE)
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
func (engine Qwant) Name() ProviderName {
	return ProviderQwant
}

// Inner searcher with pagination
func (engine Qwant) search(query string, offset int, count int) (qwantResults, error) {
	if count > 10 || count < 1 {
		return qwantResults{}, fmt.Errorf("incorrect count %d, expect value [1...10]", count)
	}
	if offset > 40 || offset < 0 {
		return qwantResults{}, fmt.Errorf("incorrect offset %d, expect value [0...40]", offset)
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
	results := qwantResults{}
	if err := helpers.RequestJSON(&results, u); err != nil {
		return qwantResults{}, err
	}

	// Handling Qwant errors
	switch results.Data.ErrorCode {
	case 0: //< All is ok
	case 429:
		return qwantResults{}, errors.NewIPBanned(fmt.Errorf("%d", results.Data.ErrorCode))
	case 14:
		return qwantResults{}, errors.NewBadRequestError(fmt.Errorf("%d", results.Data.ErrorCode))
	default:
		return qwantResults{}, fmt.Errorf("%d", results.Data.ErrorCode)
	}

	// Checks Qwant inner error
	if results.Data.ErrorCode != 0 {
		return qwantResults{}, fmt.Errorf("qwant error: %d", results.Data.ErrorCode)
	}

	return results, nil
}

// Results of Qwant search
type qwantResults struct {
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
