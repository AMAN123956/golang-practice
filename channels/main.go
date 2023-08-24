// channels are way through which my multiple go routines
// can talk with each other or interact with each other

// channels only allow values if prior some one is listening to the channel
// for eg.

// 1. myCh <- 5
//    access value from channel
//    fmt.Println(<-myCh)
// above code will not work as no goroutine is listening to channel prior before channel is sending value

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GoLang")
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{} // create a pointer

	wg.Add(2)
	// to make unidirectional
	// ch <-chan receives
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)

		// to check if channel is open or closed
		// val, isChannelOpen := <-ch
		// fmt.Println(isChannelOpen)
		// fmt.Println(val)
		wg.Done()
	}(myCh, wg)
	// chan<- sends
	go func(ch chan int, wg *sync.WaitGroup) {
		// ch <- 5
		myCh <- 6
		// close(ch)
		// for two channel values we need two listener println statements
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
