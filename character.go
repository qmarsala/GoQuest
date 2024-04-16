package main

type Background struct {
	Name        string
	Description string
	SkillBoosts map[string]int
}

type ActiveSlot struct {
	Name interface{}
	Item interface{}
}

type InventorySlot struct {
	Name     interface{}
	Item     interface{}
	Quantity int
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
	Inventory   [28]InventorySlot
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
		Inventory:   [28]InventorySlot{},
	}
	return newCharacter
}
