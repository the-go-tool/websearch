package provider

import "testing"

func TestNewUnofficialDuckDuckGo(t *testing.T) {
	uduckduckgo := NewUnofficialDuckDuckGo()
	res, err := uduckduckgo.Search("test", 65)
	if err != nil {
		t.Errorf("search error: %s", err)
	}
	if len(res) != 65 {
		t.Errorf("incorrect results count, expect 65, got %d", len(res))
	}
	for _, item := range res {
		if len(item.Title) == 0 {
			t.Errorf("empty title")
		}
	}
}
