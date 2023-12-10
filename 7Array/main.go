package main

import "fmt"

func main() {
	fmt.Println("Wellcome to arrays in Golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "mango"

	fmt.Println("Fruit List is :", fruitList)
	fmt.Println("Fruit List is :", len(fruitList))

	var vegList = [3]string{"potato", "Bean", "Mushroom"}
	fmt.Println("Veg List is :", len(vegList))

}
