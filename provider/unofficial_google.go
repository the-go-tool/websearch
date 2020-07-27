package provider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/url"
	"strconv"
	"websearch/helpers"
)

// The unofficial google provider name
const ProviderUnofficialGoogle = ProviderName("unofficial_google")

// The unofficial google web search provider
type UnofficialGoogle struct {
	api url.URL
}

// The config for unofficial google provider
type UnofficialGoogleConfig struct{}

// Makes a new unofficial google web search provider
func NewUnofficialGoogle(config ...UnofficialGoogleConfig) UnofficialGoogle {
	api := url.URL{
		Scheme: "https",
		Host:   "google.com",
		Path:   "/search",
	}

	return UnofficialGoogle{
		api: api,
	}
}

// Makes web search
func (engine UnofficialGoogle) Search(query string, count int) (Results, error) {
	results := make(Results, 0, count)

	u := engine.api
	params := map[string]string{
		"q": query,
	}

	// Fetching results
	for i := 0; i < count; i++ {
		params["start"] = strconv.Itoa(i * 10)
		u.RawQuery = helpers.ParamsRender(params)

		doc, err := helpers.RequestHTML("GET", u)
		if err != nil {
			return nil, err
		}

		fmt.Println("KJHGFDGHJ")
		doc.Find("#res #rso .g").Map(func(i int, selection *goquery.Selection) string {
			fmt.Println(selection.Text())
			return ""
		})
	}

	return results[:count], nil
}

// Returns provider name
func (engine UnofficialGoogle) Name() ProviderName {
	return ProviderUnofficialGoogle
}
