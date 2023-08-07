package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "A sample text file"

	file, err := os.Create("./test.txt")

	if err != nil {
		panic(err)
	}

	result, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	defer file.Close()
	readFile("./test.txt")
}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}
	fmt.Println(dataByte)
	// parsing byte data to string
	fmt.Println(string(dataByte))
}
