package main

type Ability struct {
	Name        string
	Description string
	Score       int
	Modifier    int
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
