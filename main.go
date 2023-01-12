package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
	var matchInfo []string

	for fileScanner.Scan() {
		line := fileScanner.Text()

		match, err := regexp.MatchString("info,team", line)
		if err != nil {
			log.Fatal("Error Matching Line", line, err)
		}
		if match {
			matchInfo = append(matchInfo, line)
			if len(matchInfo) == 2 {
				break
			}
		}
	}

	fmt.Println(len(matchInfo))
	fmt.Println(matchInfo)

	readFile.Close()
	//matchData := loadMatch(matchFile)

	//fmt.Println("Here is the raw match data", data)
	//printMatchInfo(matchData)
	// printMatchWinner()
	// printMatch11s()
}
