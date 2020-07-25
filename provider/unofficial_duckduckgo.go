package provider

import (
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"websearch/helpers"
)

// The unofficial DuckDuckGo provider name
const ProviderUnofficialDuckDuckGo = ProviderName("unofficial_duckduckgo")

// The unofficial DuckDuckGo web search provider
type UnofficialDuckDuckGo struct {
	api url.URL
}

// The config for unofficial DuckDuckGo provider
type UnofficialDuckDuckGoConfig struct{}

// Makes a new unofficial DuckDuckGo web search provider
func NewUnofficialDuckDuckGo(config ...UnofficialDuckDuckGoConfig) UnofficialDuckDuckGo {
	api := url.URL{
		Scheme: "https",
		Host:   "html.duckduckgo.com",
		Path:   "/html/",
	}

	return UnofficialDuckDuckGo{
		api: api,
	}
}

// Makes web search
func (engine UnofficialDuckDuckGo) Search(query string, count int) (Results, error) {
	results := make(Results, 0, count)

	var res Results
	var err error
	var paramsNext map[string]string

	// Initial request with first page
	res, paramsNext, err = engine.nextSearch(map[string]string{
		"q": query,
	})
	if err != nil {
		return nil, err
	}
	results = append(results, res...)

	// Next page results cycle
	for {
		if len(results) >= count {
			break
		}
		res, paramsNext, err = engine.nextSearch(paramsNext)
		if err != nil {
			return nil, err
		}
		results = append(results, res...)
	}

	return results[:count], nil
}

// Returns provider name
func (engine UnofficialDuckDuckGo) Name() ProviderName {
	return ProviderUnofficialDuckDuckGo
}

func (engine UnofficialDuckDuckGo) nextSearch(form map[string]string) (Results, map[string]string, error) {
	api := engine.api
	api.RawQuery = helpers.ParamsRender(form)

	// Gets response
	doc, err := helpers.RequestHTML("POST", api)
	if err != nil {
		return nil, nil, err
	}

	// Fetching results
	docResults := doc.Find("#links.results .result")
	results := make(Results, 0, docResults.Length())
	docResults.Map(func(i int, selection *goquery.Selection) string {
		title := selection.Find(".result__title a").Text()
		link, _ := selection.Find(".result__title a").Attr("href")
		desc := selection.Find(".result__snippet").Text()

		u, _ := url.Parse(link)
		results = append(results, Result{
			Title:       title,
			Description: desc,
			Link:        *u,
		})

		return ""
	})

	// Fetching next page params
	paramsNext := map[string]string{}
	navLinks := doc.Find(".nav-link")
	navLink := navLinks.Get(0)
	if navLinks.Length() == 2 {
		navLink = navLinks.Get(1)
	}
	inputs := goquery.NewDocumentFromNode(navLink).Find("form input")
	inputs.Map(func(i int, selection *goquery.Selection) string {
		name, _ := selection.Attr("name")
		value, _ := selection.Attr("value")
		if len(name) > 0 {
			paramsNext[name] = value
		}
		return ""
	})

	return results, paramsNext, nil
}
