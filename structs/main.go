package main

import "fmt"

type Person struct {
	Name string // Exported field
	age  int    // Unexported field
}

func main() {
	p := Person{
		Name: "John",
		age:  30,
	}

	fmt.Println("Name:", p.Name) // Accessing the exported field - OK
	// fmt.Println("Age:", p.age) // Attempting to access the unexported field - Error
}

// To Fix It
// package main

// import "fmt"

// type Person struct {
//     Name string // Exported field
//     age  int    // Unexported field
// }

// func main() {
//     p := Person{
//         Name: "John",
//         age:  30,
//     }

//     fmt.Println("Name:", p.Name)
//     fmt.Println("Age:", p.getAge()) // Accessing the unexported field via a method
// }

// // getAge is a method defined for the Person struct to access the unexported field "age"
// func (p Person) getAge() int {
//     return p.age
// }
