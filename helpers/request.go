package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Makes HTTP request, handles errs and parses result
func RequestJSON(result interface{}, url url.URL) error {
	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}

	// If status code isn't 2xx
	if resp.StatusCode/100 != http.StatusOK/100 {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Answer reading
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Result parsing
	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	return nil
}

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
