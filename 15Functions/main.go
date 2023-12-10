package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in Golang")
	greeter()

	result := adder(3, 5)
	prores, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro Result is :", prores)
	fmt.Println("Pro Message is :", myMessage)

	fmt.Println("Result is:", result)
}

func greeter() {
	fmt.Println("Namastee from Golang")
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi this is my Pro Result"
}

func adder(valOne int, valTwo int) int { //We also need to mention the type of the value we are returning
	return valOne + valTwo
}
