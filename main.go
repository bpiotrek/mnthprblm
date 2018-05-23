package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	stayStrat := func(a, b *int) *int { return a }
	changeStrat := func(a, b *int) *int { return b }

	stayPlayerAP1 := NewPlayer("stayPlayer", stayStrat)
	changePlayerAP1 := NewPlayer("changePlayer", changeStrat)

	stayPlayerAP2 := NewPlayer("stayPlayer", stayStrat)
	changePlayerAP2 := NewPlayer("changePlayer", changeStrat)

	for i := 0; i < 1000000; i++ {
		SimulateGame(r,
			stayPlayerAP1, changePlayerAP1,
			stayPlayerAP2, changePlayerAP2,
		)
	}

	fmt.Printf(
		"First approach: same game\n%v\n%v\n\nSecond approach: different games\n%v\n%v\n\n",
		stayPlayerAP1, changePlayerAP1,
		stayPlayerAP2, changePlayerAP2,
	)
}

func SimulateGame(r *rand.Rand, p1, p2, p3, p4 *Player) {
	// first approach
	game := NewGame(r)
	game.Play(p1)
	game.Play(p2)

	// second approach
	NewGame(r).Play(p3)
	NewGame(r).Play(p4)
}
