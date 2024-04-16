package main

import (
	"encoding/json"
	"os"
)

const gamestateSaveFilePath string = "./data/gamestate.json"

type Phase int

const (
	open     Phase = 0
	quest    Phase = 1
	downtime Phase = 2
)

type GameState struct {
	Character Character
	Phase     Phase
}

func load() GameState {
	jsonDat, _ := os.ReadFile(gamestateSaveFilePath)
	var gamestate GameState
	json.Unmarshal(jsonDat, &gamestate)
	return gamestate
}

func save(gamestate GameState) {
	jsonDat, _ := json.Marshal(gamestate)
	os.WriteFile(gamestateSaveFilePath, jsonDat, 0644)
}
