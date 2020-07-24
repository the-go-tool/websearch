package websearch

import (
	"errors"
	"fmt"
	"websearch/provider"
	"websearch/provider/errs"
)

func ExampleNew_ErrorHandling() {
	web := New(provider.NewUnofficialQwant())

	res, err := web.Search("test", 25)
	if err != nil {
		if errors.As(err, &errs.IPBannedError{}) {
			fmt.Println("your are banned by IP")
		}
		panic(err)
	}

	fmt.Println(res)
	// [{
	//		Title: string,
	//		Description: string,
	//		Link: url.URL,
	//		Provider: string,
	// },...]
}
