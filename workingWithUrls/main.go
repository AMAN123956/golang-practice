package main

import (
	"fmt"
	"net/url"
)

const testUrl string = "https://aman123956.github.io/aman?name=Aman&role=Web"

func main() {
	result, _ := url.Parse(testUrl)
	fmt.Println(result.Scheme) // https
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery) // list down parameter list  name=Aman&role=Web

	// key value pair
	qparam := result.Query()
	fmt.Println(qparam)         // map[name:[Aman] role:[Web]]
	fmt.Println(qparam["name"]) // [Aman]

	queryValues := url.Values{}
	queryValues.Add("user", "aman")
	// Construction a url
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "aman123956.github.io",
		Path:     "/test",
		RawQuery: queryValues.Encode(),
	}

	anotherUrl := partsOfUrl.String()

	fmt.Println(anotherUrl) // https://aman123956.github.io/test?user=aman

}
