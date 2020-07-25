package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Makes HTTP request, handles errs and returns raw result
func Request(method string, url url.URL, headers map[string]string) ([]byte, error) {
	// Custom request
	req, err := http.NewRequest(method, url.String(), strings.NewReader("q=test&b="))
	if err != nil {
		return nil, err
	}

	// Custom headers
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	// Makes request
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// If status code isn't 2xx
	if resp.StatusCode/100 != http.StatusOK/100 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Answer reading
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Makes HTTP request and parses result as JSON
func RequestJSON(result interface{}, url url.URL) error {
	data, err := Request("GET", url, nil)
	if err != nil {
		return err
	}

	// Result parsing
	if err := json.Unmarshal(data, result); err != nil {
		return err
	}

	return nil
}

// Makes HTTP request and parses result as HTML
func RequestHTML(method string, url url.URL) (*goquery.Document, error) {
	headers := map[string]string{
		"accept":     "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
		"user-agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/83.0.4103.61 Chrome/83.0.4103.61 Safari/537.36",
	}

	data, err := Request(method, url, headers)
	if err != nil {
		return nil, err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return doc, nil
}
