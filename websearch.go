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
func (webSearch WebSearch) Search(query string, maxCount ...int) (provider.Results, error) {
	results, err := webSearch.provider.Search(query, maxCount...)
	if err != nil {
		return nil, NewError(err)
	}

	// Marks provider
	for i := range results {
		results[i].Providers = []provider.ProviderName{
			webSearch.provider.Name(),
		}
	}

	return results, nil
}
