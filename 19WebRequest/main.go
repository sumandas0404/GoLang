package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response is of the type: %T \n", response)

	defer response.Body.Close() //It's The caller responsibility to close the connection request

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
