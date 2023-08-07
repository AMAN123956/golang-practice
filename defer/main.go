package main

import "fmt"

func main() {
	defer fmt.Println("World1")
	defer fmt.Println("World2")
	defer fmt.Println("World3")
	fmt.Println("Hello")
	myDefer()
}

// Defer code is executed at last
// functions declared with defer are executed in LIFO Order

// output
// Hello
// World3
// World2
// World1

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//Output
// Hello
// 4 3 2 1 0
// World3
// World2
// World1
