package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	gamestate := GameState{}
inputLoop:
	for {
		choice := readString("1) New Game\n2) Load Game\nQ) Quit\n")

		switch strings.ToLower(choice) {
		case "1":
			name := readString("Name:")
			description := readString("Description:")
			motivation := readString("Motivation:")
			catalyst := readString("Catalyst:")
			c := createNewCharacter(name, description, motivation, catalyst)
			gamestate = GameState{Character: c}
			save(gamestate)
			break inputLoop
		case "2":
			gamestate = load()
			break inputLoop
		case "q":
			os.Exit(0)
		}
	}

	fmt.Println(gamestate)
}

func readString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
