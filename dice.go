package main

import "math/rand/v2"

type Dice struct {
	Sides int
}

func (d Dice) Roll() int {
	return rand.IntN(d.Sides) + 1
}

func (d Dice) RollN(dieCount int) (total int, rolls []int) {
	for i := 0; i < dieCount; i++ {
		roll := d.Roll()
		rolls = append(rolls, roll)
		total += roll
	}
	return
}
