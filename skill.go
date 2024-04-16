package main

type Skill struct {
	Name        string
	Description string
	Score       int
	Modifier    int
	Used        bool
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
