package main

type Background struct {
	Name        string
	Description string
	SkillBoosts map[string]int
}

type Item struct {
	Name     string
	Quantity int
}

type ActiveSlot struct {
	Type string
	Item Item
}

type Character struct {
	Name        string
	Description string
	Motivation  string
	Catalyst    string
	Background  Background
	Skills      map[string]Skill
	Abilities   map[string]Ability
	ActiveSlots [6]ActiveSlot
	Inventory   [28]Item
}

func createNewCharacter(name string, description string, motivation string, catalyst string) Character {
	newCharacter := Character{
		Name:        name,
		Description: description,
		Motivation:  motivation,
		Catalyst:    catalyst,
		Background:  Background{},
		Skills:      defaultSkills,
		Abilities:   defaultAbilities,
		ActiveSlots: [6]ActiveSlot{},
		Inventory:   [28]Item{},
	}
	return newCharacter
}
