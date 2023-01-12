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
	info       Info
	playing11s []string
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
