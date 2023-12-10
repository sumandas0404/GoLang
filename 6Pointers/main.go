package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	// var ptr *int
	// fmt.Println("value of Pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of Actual Ponter is ", ptr)  //stores address (Reference)
	fmt.Println("Value of Actual Ponter is ", *ptr) //stores the value

	*ptr = *ptr * 2
	fmt.Println("New value is:", myNumber)

}
