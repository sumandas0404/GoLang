package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thrusday", "Friday", "saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {      //1)
	// 	fmt.Println(days[d])
	// }

	// for i := range days {                 //2)
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Println("Index is %v and the value is %v\n", index, day)
	}
	//OR
	for _, day := range days {
		fmt.Println("Index is and the value is %v\n", day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 5 {
			break
		}
		if rougueValue == 6 {
			rougueValue++
			continue
		}

		if rougueValue == 2 { //GOTO STATEMENT
			goto lco
		}

		fmt.Println("value is: ", rougueValue)
		rougueValue++ //++rougueValue is not applicable in GOLANG
	}

	//GOTO STATEMENT
lco:
	fmt.Println("Jumping at this location")
}
