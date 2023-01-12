package match

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Match struct {
	teams           []string
	venue           string
	winner          string
	player_of_match string
	playing11s      []string
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
			m.teams = append(m.teams, infoline[2])
			if len(m.teams) == 2 {
				break
			}
		}
	}
	fmt.Println("Match between teams ", m.teams)
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
			m.venue = venueline[2]
			break
		}
	}
	fmt.Println("Match at Venue ", m.venue)

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
			m.winner = line[2]
			break
		}
	}
	fmt.Println("Match Winner ", m.winner)
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
			m.player_of_match = line[2]
			break
		}
	}
	fmt.Println("Player of the match ", m.player_of_match)
}

func (m *Match) AddPlaying11s(matchFile string) {
	readFile, err := os.Open(matchFile)
	defer readFile.Close()

	if err != nil {
		fmt.Println(err)
	}
	matchFileScanner := bufio.NewScanner(readFile)
	matchFileScanner.Split(bufio.ScanLines)

	for matchFileScanner.Scan() {
		line := matchFileScanner.Text()

		match, err := regexp.MatchString("info,player,", line)
		if err != nil {
			log.Fatal("Error matching for team info", line, err)
		}
		if match {
			infoline := strings.Split(line, ",")
			m.playing11s = append(m.playing11s, infoline[3])
			if len(m.playing11s) == 22 {
				break
			}
		}
	}
	fmt.Println("Match playing 11s ", m.playing11s)
}
