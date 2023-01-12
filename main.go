package main

import (
	"cric/match"
	"fmt"
)

func main() {
	defer fmt.Println("Goodbye Wolf, come back.")
	fmt.Println("Hello Wolf")

	matchPath := "./data/odis_male_csv/"
	matchName := "1157378.csv"
	matchFile := fmt.Sprintf("%v%v", matchPath, matchName)

	m := &match.Match{}

	m.AddTeams(matchFile)
	m.AddVenue(matchFile)
	m.AddWinner(matchFile)
	// printMatch11s()

	fmt.Println("Here's the match", m)
}
