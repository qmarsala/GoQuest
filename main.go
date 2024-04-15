package main

import "fmt"

func main() {
	fmt.Println("hello world")
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
