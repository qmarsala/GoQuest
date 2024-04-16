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
	Used        bool
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

var defaultAbilities = map[string]Ability{
	"Agility": {
		Name:        "Agility",
		Description: "",
		Score:       1,
		Modifier:    0,
	},
	"Strength": {
		Name:        "Strength",
		Description: "",
		Score:       1,
		Modifier:    0,
	},
	"Intellect": {
		Name:        "Intellect",
		Description: "",
		Score:       1,
		Modifier:    0,
	},
}

var defaultSkills = map[string]Skill{
	"Archaeology": {
		Name:        "Archaeology",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Attack": {
		Name:        "Attack",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Cooking": {
		Name:        "Cooking",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Crafting": {
		Name:        "Crafting",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Defense": {
		Name:        "Defense",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Dungeoneering": {
		Name:        "Dungeoneering",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Farming": {
		Name:        "Farming",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Fishing": {
		Name:        "Fishing",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Fletching": {
		Name:        "Fletching",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Herblore": {
		Name:        "Herblore",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Hunting": {
		Name:        "Hunting",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Magic": {
		Name:        "Magic",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Mining": {
		Name:        "Mining",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Prayer": {
		Name:        "Prayer",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Ranged": {
		Name:        "Ranged",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Runecrafting": {
		Name:        "Runecrafting",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Slayer": {
		Name:        "Slayer",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Smithing": {
		Name:        "Smithing",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Summoning": {
		Name:        "Summoning",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Thieving": {
		Name:        "Thieving",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
	"Woodcutting": {
		Name:        "Woodcutting",
		Description: "",
		Score:       1,
		Modifier:    0,
		Used:        false,
	},
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

func CreateNewCharacter(name string, description string, motivation string, catalyst string) Character {
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
