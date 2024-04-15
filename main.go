package main

import "fmt"

func main() {
	p := createNewCharacter()
	fmt.Println("hello", p.Name)
}

func createNewCharacter() Player {
	fmt.Println("Name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Description:")
	var description string
	fmt.Scanln(&description)

	fmt.Println("Motivation:")
	var motivation string
	fmt.Scanln(&motivation)

	fmt.Println("Catalyst:")
	var catalyst string
	fmt.Scanln(&catalyst)

	p := Player{
		Name:        name,
		Description: description,
		Motivation:  motivation,
		Catalyst:    catalyst,
		Skills:      [28]Skill{},
		Abilities:   [3]Ability{},
	}
	return p
}

type Ability struct {
	Name        string
	Description string
	Score       int
	Modifier    int
}

type Skill struct {
	Name        string
	Description string
	Score       int
	Modifier    int
}

type Background struct {
	Name        string
	Description string
	SkillBoosts map[string]int
}

type Player struct {
	Name        string
	Description string
	Motivation  string
	Catalyst    string
	Skills      [28]Skill
	Abilities   [3]Ability
}
