[![GitHub](https://img.shields.io/github/license/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/blob/master/LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/the-go-tools/websearch)](https://blog.golang.org/go1.13)
[![Go Report Card](https://goreportcard.com/badge/github.com/the-go-tools/websearch)](https://goreportcard.com/report/github.com/the-go-tools/websearch)
[![code-coverage](http://gocover.io/_badge/github.com/the-go-tools/websearch)](https://gocover.io/github.com/the-go-tools/websearch)
[![GoDoc](https://godoc.org/github.com/the-go-tools/websearch?status.svg)](https://godoc.org/github.com/the-go-tools/websearch)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/releases)
[![GitHub last commit](https://img.shields.io/github/last-commit/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/commits/master)
[![GitHub issues](https://img.shields.io/github/issues/the-go-tools/websearch)](https://github.com/the-go-tools/websearch/issues)

# The Go Tools :: Web Search
> :construction: The tool is in under construction and the
> API can change

The simple tool to use any web search engines like Google, Yandex, Bing,
Qwant, DuckDuckGo and so on.

Supports now:
- [X] Unofficial Qwant
- [ ] Qwant
- [ ] Unofficial DuckDuckGo
- [ ] DuckDuckGo
- [ ] More: Google, Yandex, Bing, Yahoo etc

## :fast_forward: Fast Start

### :arrow_down: Get It
> go get github.com/the-go-tools/websearch

Then add imports:
```go
import (
    "github.com/the-go-tools/websearch"
    "github.com/the-go-tools/websearch/provider"
)
```

### :pencil: Configure It
```go
web := websearch.New(provider.NewQwant("en_US"))
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

## :question: Q/A

> **Q:** Should I use **unofficial** providers?  
> **A:** Maybe. It depends on stability you expect.
> Official APIs require they token and may take taxes.
> Unofficial APIs are free, but they are unstable and your
> IP may be banned for several minutes.
> So, if you have your personal/home project or you
> don't want pay then choose unofficial.

### :star: Please, star it if you find it helpful

### :link: Similar projects
If this project doesn't fit.
- :link: https://github.com/serpapi/google-search-results-golang
- :link: https://github.com/rocketlaunchr/google-search
- :link: https://github.com/schollz/googleit
