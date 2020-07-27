package provider

import "testing"

func TestNewUnofficialGoogle(t *testing.T) {
	ugoogle := NewUnofficialGoogle()
	res, err := ugoogle.Search("test", 25)
	if err != nil {
		t.Errorf("search error: %s", err)
	}
	if len(res) != 25 {
		t.Errorf("incorrect results count, expect 25, got %d", len(res))
	}
	for _, item := range res {
		if len(item.Title) == 0 {
			t.Errorf("empty title")
		}
	}
}
