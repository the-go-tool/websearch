package websearch

import (
	"testing"
	"websearch/provider"
)

func TestWebSearch_Search(t *testing.T) {
	web := New(provider.NewQwant("en_US"))
	res, err := web.Search("test", 25)
	if err != nil {
		t.Fatalf("search error: %s", err)
	}
	if len(res) != 25 {
		t.Fatalf("incorrect results count, expect 25, got %d", len(res))
	}
	for _, item := range res {
		if len(item.Title) == 0 {
			t.Fatalf("empty title")
		}
	}
}
