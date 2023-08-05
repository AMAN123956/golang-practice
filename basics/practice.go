package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const PI = 3.14 // constant declaration

func main() {
	var age int = 21 // variable declaration
	x := "Dixit"     //  short variable declaration
	fmt.Println("Hello World Aman", x, age, PI)

	// Taking Input from User
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pls Enter Your value")
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered Rating is", input)

	// Type of input is string

	// Type Conversion in goLang

	// We cannot add string to int in go

	// we need to implement type conversion for this we use a package named strconv
	numInput, err := strconv.ParseFloat(strings.TrimSpace(input), 32)

	if err != nil {
		fmt.Println(err)
	}

	// strings package used to trim space has many other built-in functions like
	// TrimSpace,TrimLeft

	res := numInput + 1
	fmt.Println("Your answer is", res)

	// Time
	fmt.Println("Time study in go lang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// formatting time
	fmt.Println(presentTime.Format("01-02-2006")) // for formatting in terms of date -  month - year
	// this value of 01-02-2006 is fixed

	// for day
	fmt.Println(presentTime.Format("01-02-2006 Monday"))
	// for time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// create a slice
	// myList := make([]int,5,10)
	// myList = append(myList, 1, 2, 3, 4, 5, 6, 7, 8)
	// fmt.Println("List is", myList)
	// In this first 5 index are filled as zero
	// as make function initially creates a slice of length 5 and initializes its value with 0

	// To remove Zeros from start
	// myList := make([]int,0,10)
	var myList = []string{"a", "b", "c", "d"} // initialization is must
	myList = append(myList, "e", "f")
	fmt.Println("Improved Slice", myList)

	// Another way of creating a slice
	// var keys []string

	// Difference between array and slice in golang
	// 1. No need of declaring size in slice while in case of array it is must

	// Remove a value from slice based on index
	var courses = []string{"a", "b", "c", "d", "e", "f"}

	// we need to remove index = 2
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Updated slice after 2 index removal", courses) // Updated slice after 2 index removal [a b d e f] -- output

	// Creating Maps in Golang
	languages := make(map[string]string) // key,value
	languages["a"] = "a"
	languages["b"] = "b"
	languages["c"] = "c"
	fmt.Println("map is", languages) // map is map[a:a b:b c:c]

	// To Loop in key,value in map
	for key, value := range languages {
		fmt.Println("key", key, value)
	}
}
