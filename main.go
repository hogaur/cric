package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	defer fmt.Println("Goodbye Wolf, come back.")
	fmt.Println("Hello Wolf")

	matchPath := "./data/odis_male_csv/"
	matchName := "1157378.csv"
	matchFile := fmt.Sprintf("%v%v", matchPath, matchName)

	readFile, err := os.Open(matchFile)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	m := &Match{}

	m.addTeams(*fileScanner)
	//printMatchInfo(matchData)
	// printMatchWinner()
	// printMatch11s()

	fmt.Println("Here's the match", m)

	defer readFile.Close()
}

type Match struct {
	teams []string
}

func (m *Match) addTeams(matchFileScanner bufio.Scanner) {
	for matchFileScanner.Scan() {
		line := matchFileScanner.Text()

		match, err := regexp.MatchString("info,team", line)
		if err != nil {
			log.Fatal("Error matching for team info", line, err)
		}
		if match {
			infoline := strings.Split(line, ",")
			m.teams = append(m.teams, infoline[2])
			if len(m.teams) == 2 {
				break
			}
		}
	}
}
