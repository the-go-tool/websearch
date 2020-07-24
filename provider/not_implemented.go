package provider

import (
	"errors"
	"websearch/provider/errs"
)

// The not implemented provider name
const ProviderNotImplemented = ProviderName("not_implemented")

// The not implemented web search provider
type NotImplemented struct{}

// The config for not implemented provider
type NotImplementedConfig struct{}

// Makes a new not implemented web search provider
func NewNotImplemented(config ...NotImplementedConfig) NotImplemented {
	return NotImplemented{}
}

// Makes web search
func (engine NotImplemented) Search(query string, count int) (Results, error) {
	return Results{}, errs.NewNotImplemented(errors.New("don't use not implemented provider"))
}

// Returns provider name
func (engine NotImplemented) Name() ProviderName {
	return ProviderNotImplemented
}
