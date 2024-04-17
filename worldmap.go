package main

import (
	"fmt"
	"math"
)

type Location struct {
	Name        string
	Description string
}

type Path struct {
	Distance    int
	Destination Location
}

type World struct {
	Locations map[string][]Path
}

func (w *World) addPath(src Location, dest Location, distance int) {
	if _, exists := w.Locations[src.Name]; !exists {
		w.Locations[src.Name] = []Path{}
	}
	if _, exists := w.Locations[dest.Name]; !exists {
		w.Locations[dest.Name] = []Path{}
	}
	w.Locations[src.Name] = append(w.Locations[src.Name], Path{Destination: dest, Distance: distance})
}

func (w *World) Print() {
	for location, paths := range w.Locations {
		for _, path := range paths {
			fmt.Printf("%s -> %s (Distance: %d)\n", location, path.Destination.Name, path.Distance)
		}
	}
}

func (w *World) CalculateDistance(src string, target string) int {
	dist := make(map[string]int)
	visited := make(map[string]bool)

	for location := range w.Locations {
		dist[location] = math.MaxInt32
	}
	dist[src] = 0

	for len(visited) < len(w.Locations) {
		minLocation := ""
		minDist := math.MaxInt32
		for location, d := range dist {
			if !visited[location] && d < minDist {
				minDist = d
				minLocation = location
			}
		}

		if minLocation == "" {
			break
		}

		visited[minLocation] = true

		for _, path := range w.Locations[minLocation] {
			location := path.Destination.Name
			if !visited[location] {
				newDist := dist[minLocation] + path.Distance
				if newDist < dist[location] {
					dist[location] = newDist
				}
			}
		}
	}
	return dist[target]
}

func initializeWorld() World {
	world := World{
		Locations: make(map[string][]Path),
	}
	lumbridge := Location{
		Name:        "Lumbridge",
		Description: "The breadbasket of Gielinor",
	}
	varrock := Location{
		Name:        "Varrock",
		Description: "The oldest and largest city in Gielinor",
	}
	falador := Location{
		Name:        "Falador",
		Description: "Home of the White Knights",
	}
	world.addPath(lumbridge, varrock, 3)
	world.addPath(lumbridge, falador, 4)
	world.addPath(varrock, lumbridge, 3)
	world.addPath(varrock, falador, 2)
	return world
}
