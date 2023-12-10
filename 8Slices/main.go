package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices")

	var fruitList = []string{"Apple", "Tomato", "peach"}
	fmt.Println("Type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana") //To add more items
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:]) //Used for slicing
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 456
	highScores[2] = 234
	highScores[3] = 689

	highScores = append(highScores, 223, 656, 344) // These way we can accomodate more values in the slice
	fmt.Println(highScores)

	sort.Ints(highScores) // sorts the values in ascending order
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores)) // returns true is values are sorted else false
}
