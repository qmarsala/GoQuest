package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	c := createNewCharacter()
	SaveCharacter(c)
}

func createNewCharacter() Character {
	name := readString("Name:")
	description := readString("Description:")
	motivation := readString("Motivation:")
	catalyst := readString("Catalyst:")

	p := Character{
		Name:        name,
		Description: description,
		Motivation:  motivation,
		Catalyst:    catalyst,
		Skills:      map[string]Skill{},
		Abilities:   map[string]Ability{},
	}
	return p
}

func readString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
