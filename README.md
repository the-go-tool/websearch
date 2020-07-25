[![GitHub](https://img.shields.io/github/license/the-go-tool/websearch)](https://github.com/the-go-tool/websearch/blob/master/LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/the-go-tool/websearch)](https://blog.golang.org/go1.13)
[![Build Status](https://travis-ci.com/the-go-tool/websearch.svg?branch=master)](https://travis-ci.com/the-go-tool/websearch)
[![Go Report Card](https://goreportcard.com/badge/github.com/the-go-tool/websearch)](https://goreportcard.com/report/github.com/the-go-tool/websearch)
[![GoDoc](https://godoc.org/github.com/the-go-tool/websearch?status.svg)](https://godoc.org/github.com/the-go-tool/websearch)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/the-go-tool/websearch)](https://github.com/the-go-tool/websearch/releases)
[![GitHub last commit](https://img.shields.io/github/last-commit/the-go-tool/websearch)](https://github.com/the-go-tool/websearch/commits/master)
[![GitHub issues](https://img.shields.io/github/issues/the-go-tool/websearch)](https://github.com/the-go-tool/websearch/issues)

# The Go Tool :: Web Search :mag_right:
> :construction: The tool is in under construction and the
> API can change

This is simple tool to use any web search engines like Google, Yandex, Bing,
Qwant, DuckDuckGo and so on.

Supports now:
- [X] Unofficial Qwant
- [X] Unofficial DuckDuckGo
- [ ] Unofficial Google
- [ ] More: Google, Yandex, Bing, Yahoo etc

I need help. If you know any way to get official Qwant or DuckDuckGo APIs,
please, make an issue.

## :fast_forward: Fast Start

### Get It
> go get github.com/the-go-tool/websearch

Then add imports:
```go
import (
    "github.com/the-go-tool/websearch"
    "github.com/the-go-tool/websearch/provider"
    "github.com/the-go-tool/websearch/provider/errs"
)
```

### Configure It
```go
web := websearch.New(provider.NewUnofficialQwant())
```

### Use It
```go
res, err := web.Search("test", 25)
if err != nil {
    // ...
}

fmt.Println(res)
// [{
//		Title: string,
//		Description: string,
//		Link: url.URL,
//		Provider: string,
// },...]
```

## :arrow_forward: More Detailed Start

### Provider Configuration
Some providers require configuration.
It can be optional or not.
If you have a token or any other credentials for official APIs,
you can pass them by provider config.
```go
web := websearch.New(provider.NewUnofficialQwant(provider.UnofficialQwantConfig{
    Locale: "ru_RU",
}))
```

### Error Handling
The library has several own errors.  
Every error in websearch wrapped into websearch.Error,
so you can handle only errors from this library like:
```go
res, err := web.Search("test", 25)
if err != nil {
    if errors.As(err, &websearch.Error{}) {
        // ...
    }
    // ...
}
```

Next, providers have common specific errors.  
You can get IP ban when use unofficial API and you can check this case so:
```go
res, err := web.Search("test", 25)
if err != nil {
    if errors.As(err, &errs.IPBannedError{}) {
        fmt.Println("your are banned by IP")
    }
    panic(err)
}
```

## :question: Q/A

> **Q:** Should I use **unofficial** providers?  
> **A:** Maybe. It depends on stability you expect.
> Official APIs require they token and may take taxes.
> Unofficial APIs are free, but they are unstable and your
> IP may be banned for several minutes.
> So, if you have your personal/home project or you
> don't want pay then choose unofficial.

### :star: Please, star it if you find it helpful

#### Similar projects
If this project doesn't fit.
- :link: https://github.com/serpapi/google-search-results-golang
- :link: https://github.com/rocketlaunchr/google-search
- :link: https://github.com/schollz/googleit
