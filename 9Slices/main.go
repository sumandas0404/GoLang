package main

import "fmt"

func main() {
	var courses = []string{"Reactjs", "javaScript", "Swift", "Python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
