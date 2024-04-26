package main

type Ability struct {
	Name        string
	Description string
	Score       int
	Modifier    int
}

var defaultAbilities = []Ability{
	{
		Name:        "Agility",
		Description: "",
		Score:       1,
		Modifier:    0,
	},
	{
		Name:        "Strength",
		Description: "",
		Score:       1,
		Modifier:    0,
	},
	{
		Name:        "Intellect",
		Description: "",
		Score:       1,
		Modifier:    0,
	},
}
