package main

import (
	"cric/match"
	"fmt"
)

func main() {
	defer fmt.Println("Goodbye Wolf, come back.")
	fmt.Println("Hello Wolf")

	matchPath := "./data/odis_male_csv/"
	matchFileName := "1157378.csv"
	matchFilePath := fmt.Sprintf("%v%v", matchPath, matchFileName)

	m := &match.Match{}

	m.AddTeams(matchFilePath)
	m.AddVenue(matchFilePath)
	m.AddWinner(matchFilePath)
	m.AddPlayerOfMatch(matchFilePath)
	m.AddPlaying11s(matchFilePath)

	fmt.Println("Here's the match", m)
}
