package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	d6 := Dice{6}
	total, rolls := d6.RollN(3)
	fmt.Println(total)
	fmt.Println(rolls)

	d4 := Dice{4}
	roll := d4.Roll()
	fmt.Println(roll)
	// name := readString("Name:")
	// description := readString("Description:")
	// motivation := readString("Motivation:")
	// catalyst := readString("Catalyst:")
	// c := CreateNewCharacter(name, description, motivation, catalyst)
	// SaveCharacter(c)
}

func readString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}
