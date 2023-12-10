package main

import "fmt"

func main() {
	defer fmt.Println("World") //writing defer takes the statement and moves it after the curly braces.
	defer fmt.Println("One")   // defer executes in last in fast out manner.
	defer fmt.Println("Two")   // so first Two then one then World will be printed
	fmt.Println("Hello")
	mydefer()
}

func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// world, one , two
//0,1,2,3,4
//PRINT: hello, 43210, two, one , world
