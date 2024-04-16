package main

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func loadCharacter() {
	var character Character
	jsonDat, _ := os.ReadFile(characterSaveFilePath)
	json.Unmarshal(jsonDat, &character)
	fmt.Println(character)
}

func saveCharacter(character Character) {
	jsonDat, _ := json.Marshal(character)
	os.WriteFile(characterSaveFilePath, jsonDat, 0644)
}

func createNewCharacter(name string, description string, motivation string, catalyst string) Character {
	newCharacter := Character{
		Name:        name,
		Description: description,
		Motivation:  motivation,
		Catalyst:    catalyst,
		Skills:      defaultSkills,
		Abilities:   defaultAbilities,
	}
	return newCharacter
}
