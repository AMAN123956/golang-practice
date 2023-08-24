package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	fmt.Println("Race Condition")
// 	wg := &sync.WaitGroup{}
// 	var score = []int{0}

// 	wg.Add(3)

// 	// creating three go routines
// 	go func(wg *sync.WaitGroup) {
// 		fmt.Println("One R")
// 		score = append(score, 1)
// 		wg.Done()
// 	}(wg)
// 	go func(wg *sync.WaitGroup) {
// 		fmt.Println("Two R")
// 		score = append(score, 2)
// 		wg.Done()
// 	}(wg)
// 	go func(wg *sync.WaitGroup) {
// 		fmt.Println("Three R")
// 		score = append(score, 3)
// 		wg.Done()
// 	}(wg)

// 	wg.Wait()
// 	fmt.Println(score)
// }

// result can come in any order as it all depends on threads
// order is not guarenteed in above code

// Now Comes race condition logic
func main() {
	fmt.Println("Race Condition")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	var score = []int{0}

	wg.Add(3)

	// creating three go routines
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}

// to check whether race condition exist or not
// run command -  go run --race main.go
// if it gives nothing then no race condition exists

// mutex play a key role in removing race condition from the code
