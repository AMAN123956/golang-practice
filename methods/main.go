package main

import "fmt"

type Person struct {
	Name   string // Exported field
	age    int    // Unexported field
	Status bool
}

func main() {
	p1 := Person{
		Name:   "John",
		age:    30,
		Status: true,
	}

	fmt.Println("Name:", p1.Name) // Accessing the exported field - OK
	// fmt.Println("Age:", p.age) // Attempting to access the unexported field - Error
	p1.GetStatus()
}

// The function attached to class is called as method
// here p is passed as copy
// any changes made on p won't be reflected on actual Person
func (p Person) GetStatus() {
	fmt.Println("Is user active", p.Status)
}
