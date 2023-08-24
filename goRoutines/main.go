package main

import (
	"fmt"
	"net/http"
	"sync"
)

// func main() {
// 	go greeter("Hello") // it creates a separate thread for execution
// 	greeter("World")
// }

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		fmt.Println(s)
// 	}
// }

// Output is
// World -- 5 times no hello
// its bcz since first call is executing in a separate thread
// and till the time it return my main thread getx excuted and exited

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{"test"}

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://google.com",
		"https://github.com",
		"https://aman123956.github.com",
	}

	for _, web := range websiteList {
		go getStatus(web) // it is a slow process we can use go rountines to execute these in parallel
		wg.Add(1)         // it add wait group
	}

	wg.Wait() // it prevent main func from exiting untill all of it wait groups have returned
	fmt.Println(signals)
}

func getStatus(endpoint string) {
	defer wg.Done() // will be moved to end of func due to defer keyword
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Some error in endpoints")
	}

	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
}

// In case of go routines we must ensure that if
// one go routine is working at a time no other operations in the
// same memory should be performed

// this is essential when there are multiple queries writing
// in same db then this lock comes in handy to ensure that at one time only
// one write operation is performed in db
