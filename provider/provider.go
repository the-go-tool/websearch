package provider

import "net/url"

// The name of web search provider
type ProviderName string

// The unknown web search provider
const ProviderUnknown = ProviderName("unknown")

// The provider's interface
type Provider interface {
	Search(query string, maxCount ...int) (Results, error) // < The main search method with search query
	Name() ProviderName // < The provider name
}

// The web search result entries
type Results []Result

// The web search result entry
type Result struct {
	Title string
	Description string
	Link url.URL
	Providers []ProviderName
}
