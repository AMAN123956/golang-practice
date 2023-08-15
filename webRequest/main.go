package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	performPostJSONRequest()
}

func performPostJSONRequest() {
	const myUrl = "http://localhost:8002"

	// fake json payload
	requestBody := strings.NewReader(`
  {
	  "name":"Aman",
	  "price":0,
  }
  `)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	content := string(dataBytes)
	fmt.Println(content)
}
