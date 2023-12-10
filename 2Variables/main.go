package main

import "fmt"

// jwttoken := 20000 This declaration is outside the method i.e GLOBAL thus not allowed
// Constants
const LoginToken string = "Suman0404" //Public vbariable

func main() {
	var username string = "Suman"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type: %T \n", smallval)

	var smallfloat float32 = 255.565776677
	fmt.Println(smallfloat)
	fmt.Printf("variable is of type: %T \n", smallfloat)

	//default values and some aliases
	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of the type: %T \n", anotherVariable)

	//implicit type of declaration
	var website = "Suman Das"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of the type: %T \n", LoginToken)
}
