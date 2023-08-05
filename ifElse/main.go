package main

import "fmt"

func main() {
	fmt.Println("If Else")

	var result int = 10

	if result < 10 {
		fmt.Println("If")
	} else {
		fmt.Println("Else")
	}

	// special syntax
	if num := 3; num < 10 {
		fmt.Println("If")
	} else {
		fmt.Println("Else")
	}

}
