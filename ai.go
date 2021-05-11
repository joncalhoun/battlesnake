package main

import (
	"math/rand"
)

func RandomNoSuicide(g GameRequest) Move {
	moves := []Move{Up, Down, Left, Right}
	rand.Shuffle(4, func(i, j int) {
		moves[i], moves[j] = moves[j], moves[i]
	})
	best := Noop
	for _, m := range moves {
		res := MoveResult(g, m)
		switch res {
		case Fed:
			return m
		case Alive:
			best = m
		case MaybeDeath:
			if best == Noop {
				best = m
			}
		case Death:
			// noop
		}
	}
	if best == Noop {
		// Doesn't matter
		return Up
	}
	return best
}

func HungryHungryHippo(g GameRequest) Move {
	// TODO: Create the Hungry Hungry Hippo AI
	return Up
}
