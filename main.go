package main

import (
	"fmt"
	"net/http"
)

func main() {
	defer fmt.Println("Goodbye Wolf, come back.")
	fmt.Println("Hello Wolf")

	gill := "gill"

	fmt.Printf("First we get 1 player info for %v\n", gill)

	url := fmt.Sprintf("https://cricsheet.org/format/json/#players/%v", gill)
	fmt.Println("URL for", gill)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to get player info for", gill)
		fmt.Println(err)
	} else {
		fmt.Println("We received player info response for", gill)
		fmt.Println("Here is the response: ", res)
	}
}
