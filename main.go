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

	//world := initializeWorld()

gameLoop:
	for {
		// during open play
		// 	character can perform up to 3 actions (per location visited)
		// 	and any amount of travel/npc conversations
		// during quest play
		//	 character is working towards an objective
		// during downtime
		//	character can chose a used skill to level up
		//  and retry any failed actions during open play

		switch {
		case gamestate.Phase == open:
			gamestate.Phase = handleOpenPlay()
		case gamestate.Phase == quest:
			gamestate.Phase = handleQuestPlay()
		case gamestate.Phase == downtime:
			gamestate.Phase = handleDowntime()
		default:
			gamestate.Phase = handleOpenPlay()
		}
		break gameLoop
	}
	save(gamestate)
}

func handleOpenPlay() Phase {
	fmt.Println("open play")
	return quest
}

func handleQuestPlay() Phase {
	fmt.Println("questing")
	return downtime
}

func handleDowntime() Phase {
	fmt.Println("downtime")
	return open
}

func readString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
