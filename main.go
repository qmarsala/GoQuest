package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := readString("Name:")
	description := readString("Description:")
	motivation := readString("Motivation:")
	catalyst := readString("Catalyst:")
	c := CreateNewCharacter(name, description, motivation, catalyst)
	SaveCharacter(c)
}

func readString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
