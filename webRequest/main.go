package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "test_url"

func main() {
	response, err := http.Get(url) // package to make get request

	if err != nil {
		panic(err)
	}

	defer response.Body.Close() // callers responsibility to close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)

	content := string(dataBytes)
	fmt.Println(content)
}
