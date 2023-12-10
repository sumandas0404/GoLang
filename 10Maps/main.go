package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages:", languages)

	fmt.Println("JS shorts for:", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	//loops are interesting in GOLANG
	for key, value := range languages { //for each types
		fmt.Println("For key %v, value is %v \n", key, value)
	}

}
