package match

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Info struct {
	teams           []string
	venue           string
	winner          string
	player_of_match string
}

func (m *Match) AddTeams(matchFile string) {
	readFile, err := os.Open(matchFile)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	matchFileScanner := bufio.NewScanner(readFile)
	matchFileScanner.Split(bufio.ScanLines)

	for matchFileScanner.Scan() {
		line := matchFileScanner.Text()

		match, err := regexp.MatchString("info,team", line)
		if err != nil {
			log.Fatal("Error matching for team info", line, err)
		}
		if match {
			infoline := strings.Split(line, ",")
			m.info.teams = append(m.info.teams, infoline[2])
			if len(m.info.teams) == 2 {
				break
			}
		}
	}
	fmt.Println("Match between teams ", m.info.teams)
}

func (m *Match) AddVenue(matchFile string) {
	readFile, err := os.Open(matchFile)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	matchFileScanner := bufio.NewScanner(readFile)
	matchFileScanner.Split(bufio.ScanLines)

	for matchFileScanner.Scan() {
		line := matchFileScanner.Text()

		match, err := regexp.MatchString("info,venue", line)
		if err != nil {
			log.Fatal("Error matching for team info", line, err)
		}
		if match {
			venueline := strings.Split(line, ",")
			m.info.venue = venueline[2]
			break
		}
	}
	fmt.Println("Match at Venue ", m.info.venue)

}

func (m *Match) AddWinner(matchFile string) {
	readFile, err := os.Open(matchFile)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	matchFileScanner := bufio.NewScanner(readFile)
	matchFileScanner.Split(bufio.ScanLines)

	for matchFileScanner.Scan() {
		line := matchFileScanner.Text()

		match, err := regexp.MatchString("info,winner", line)
		if err != nil {
			log.Fatal("Error matching for match winner", line, err)
		}
		if match {
			line := strings.Split(line, ",")
			m.info.winner = line[2]
			break
		}
	}
	fmt.Println("Match Winner ", m.info.winner)
}
func (m *Match) AddPlayerOfMatch(matchFile string) {
	readFile, err := os.Open(matchFile)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	matchFileScanner := bufio.NewScanner(readFile)
	matchFileScanner.Split(bufio.ScanLines)

	for matchFileScanner.Scan() {
		line := matchFileScanner.Text()

		match, err := regexp.MatchString("info,player_of_match", line)
		if err != nil {
			log.Fatal("Error matching for the player of the match", line, err)
		}
		if match {
			line := strings.Split(line, ",")
			m.info.player_of_match = line[2]
			break
		}
	}
	fmt.Println("Player of the match ", m.info.player_of_match)
}
