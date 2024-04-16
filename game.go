package main

import (
	"encoding/json"
	"os"
)

const gamestateSaveFilePath string = "./gamestate.json"

type GameState struct {
	Character Character
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
