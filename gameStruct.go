package main

import "math/rand"

type Game struct {
	winningDoors int
	initialGuess int
	closedDoors  int
}

func NewGame(r *rand.Rand) *Game {
	winning := r.Intn(3)
	initial := 0
	closed := winning
	if winning == initial {
		closed = (winning + 1) % 3 // leave next doors to selected
		// closed = (winning + 2) % 3 // open next doors to selected
	}

	return &Game{
		winningDoors: winning,
		initialGuess: initial,
		closedDoors:  closed,
	}
}

func (g *Game) Play(p *Player) {
	if *p.MakeChoice(&g.initialGuess, &g.closedDoors) == g.winningDoors {
		p.Won()
	} else {
		p.Lost()
	}
}
