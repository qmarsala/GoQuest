package main

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
