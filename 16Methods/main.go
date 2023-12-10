package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	suman := User{"Suman", "sumon0404@gmail.com", true, 21}
	fmt.Println(suman)                            //prints only the values
	fmt.Printf("Suman Details are: %+v\n", suman) //prints the values along with the naming convention used here
	fmt.Printf("Name is %v and email is %v.", suman.Name, suman.Email)

	suman.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active:", u.Status)
}

func (u User) Newmail  {   //Copy is passed
	u.Email= "test@gmail.com"	//Manipulates the email
	fmt.Println("Email of this User is: ",u.Email)
}
