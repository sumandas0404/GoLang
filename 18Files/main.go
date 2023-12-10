package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "this needs to go in a file"

	file, err := os.Create("./myfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content) //For writing into the file
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is :", length)
	defer file.Close()
	readFile("./myfile.txt")

}

func readFile(filename string) { //Reading from a file
	databyte, err := ioutil.ReadFile(filename)
	// if err != nil {
	// 	panic(err)
	// }
	//OR
	checkNilErr(err)

	fmt.Println("text data inside the file is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
