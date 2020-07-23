[![GitHub](https://img.shields.io/github/license/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/blob/master/LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/the-go-tools/websearch)](https://blog.golang.org/go1.13)
[![Go Report Card](https://goreportcard.com/badge/the-go-tools/websearch)](http://goreportcard.com/report/the-go-tools/websearch)
[![code-coverage](http://gocover.io/_badge/github.com/the-go-tools/websearch)](https://gocover.io/github.com/the-go-tools/websearch)
[![GoDoc](https://godoc.org/github.com/the-go-tools/websearch?status.svg)](https://godoc.org/github.com/the-go-tools/websearch)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/releases)
[![GitHub last commit](https://img.shields.io/github/last-commit/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/commits/master)
[![GitHub issues](https://img.shields.io/github/issues/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/issues)

# The Go Tools :: Web Search
> :construction: The tool is in under construction yet

The simple tool to use any web search engines like Google, Yandex, Bing,
Qwant, DuckDuckGo and so on.

Supports now:
- [X] Qwant
- [ ] DuckDuckGo
- [ ] Google
- [ ] Yandex
- [ ] Bing
- [ ] Yahoo
- [ ] Other

## :fast_forward: Fast Start

### :arrow_down: Get It
> go get github.com/the-go-tools/websearch

### :pencil: Configure It
```go
web := New(provider.NewQwant("en_US"))
```

### :checkered_flag: Use It
```go
res, err := web.Search("test", 25)
// [
//  { 
//      Title: string, 
//      Description: string, 
//      Link: url.URL, 
//      Provider: "qwant" 
//  }, ...
// ]
```

## :arrow_forward: More Detailed Start
soon