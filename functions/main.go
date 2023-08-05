package main

import (
	"fmt"
)

func sampleFunction() {
	fmt.Println("It is a sample function")
}

func withParameters(name string, age int) {
	fmt.Println("Your Name", name)
	fmt.Println("Your Age", age)
}

// Function Return Example
func myFunction1(x int, y int) int {
	return x + y
}

// We can specify the return type and return variable name of function
func myFunction2(x int, y int) (result int) {
	result = x + y
	return result
	// also return cal also work instead of return result
}

// Multiple Return Values
func myFunction3(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func main() {
	sampleFunction()
	withParameters("Test", 18)
	fmt.Println(myFunction1(1, 2))
	fmt.Println(myFunction2(1, 2))
	fmt.Println(myFunction3(3, "Hello"))
}
