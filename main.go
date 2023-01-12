package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("Goodbye Wolf, come back.")
	fmt.Println("Hello Wolf")

	matchPath := "data/odis_male_csv"
	match := "1157378.csv"

	fmt.Printf("First we get 1 match info for %v\n", match)

	fmt.Println("Loading a match", match)
	loadMatch(matchPath, match)
}

func createMatch(data [][]string) []Match {
	return []Match{}
}

func loadMatch(csvPath, csvName string) {
	filepath := fmt.Sprintf(csvPath, "/", csvPath)
	fmt.Println("filepath for the match", filepath)

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// convert records to array of structs
	shoppingList := createMatch(data)

	// print the array
	fmt.Printf("%+v\n", shoppingList)
}

type Match struct {
}
