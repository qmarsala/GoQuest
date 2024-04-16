package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

type Character struct {
	Name        string
	Description string
	Motivation  string
	Catalyst    string
	Skills      map[string]Skill
	Abilities   map[string]Ability
}

const characterSaveFilePath string = "./character.json"

func LoadCharacter() {
	var character Character
	jsonDat, _ := os.ReadFile(characterSaveFilePath)
	json.Unmarshal(jsonDat, &character)
	fmt.Println(character)
}

func SaveCharacter(character Character) {
	jsonDat, _ := json.Marshal(character)
	os.WriteFile(characterSaveFilePath, jsonDat, 0644)
}
