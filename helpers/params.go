package helpers

import (
	"fmt"
	"net/url"
	"strings"
)

// Merges string:string sets params
func ParamsMerge(sets ...map[string]string) map[string]string {
	result := map[string]string{}
	for i := range sets {
		for key, val := range sets[i] {
			result[key] = val
		}
	}
	return result
}

// Renders params to query string
func ParamsRender(params map[string]string) string {
	parts := make([]string, 0, len(params))
	for key, val := range params {
		parts = append(parts, fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(val)))
	}
	return strings.Join(parts, "&")
}
