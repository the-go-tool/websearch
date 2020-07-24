package websearch

import (
	"websearch/provider"
)

// The main web search wrapper
type WebSearch struct {
	provider provider.Provider
}

// Makes a new web search provider
func New(provider provider.Provider) *WebSearch {
	return &WebSearch{
		provider: provider,
	}
}

// Makes web search
func (webSearch WebSearch) Search(query string, count ...int) (provider.Results, error) {
	c := 10
	if len(count) > 0 {
		c = count[0]
	}

	results, err := webSearch.provider.Search(query, c)
	if err != nil {
		return nil, NewError(err)
	}

	// Marks provider
	for i := range results {
		results[i].Provider = webSearch.provider.Name()
	}

	return results, nil
}
