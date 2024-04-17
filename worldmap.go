package main

import "fmt"

type Location struct {
	Name        string
	Description string
}

type Path struct {
	Distance    int
	Destination Location
}

type World struct {
	Locations map[Location][]Path
}

func (w *World) addPath(src Location, dest Location, distance int) {
	w.Locations[src] = append(w.Locations[src], Path{Destination: dest, Distance: distance})
}

func (w *World) Print() {
	for location, paths := range w.Locations {
		for _, path := range paths {
			fmt.Printf("%s -> %s (Distance: %d)\n", location.Name, path.Destination.Name, path.Distance)
		}
	}
}

func initializeWorld() World {
	world := World{
		Locations: make(map[Location][]Path),
	}
	lumbridge := Location{
		Name:        "Lumbridge",
		Description: "The breadbasket of Gielinor",
	}
	varrock := Location{
		Name:        "Varrock",
		Description: "The oldest and largest city in Gielinor",
	}
	falidor := Location{
		Name:        "Falador",
		Description: "Home of the White Knights",
	}
	world.addPath(lumbridge, varrock, 3)
	world.addPath(varrock, lumbridge, 3)
	world.addPath(varrock, falidor, 2)
	return world
}
