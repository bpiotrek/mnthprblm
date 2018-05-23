package main

import "fmt"

type Strategy func(*int, *int) *int

type Player struct {
	name       string
	strategy   Strategy
	totalGames int
	wonGames   int
}

func NewPlayer(name string, strategy Strategy) *Player {
	return &Player{
		name:       name,
		strategy:   strategy,
		totalGames: 0,
		wonGames:   0,
	}
}

func (p *Player) MakeChoice(initial, alternative *int) *int {
	return p.strategy(initial, alternative)
}

func (p *Player) Won() {
	p.totalGames++
	p.wonGames++
}

func (p *Player) Lost() {
	p.totalGames++
}

func (p *Player) String() string {
	total := float64(p.totalGames)
	won := float64(p.wonGames)
	lost := (total - won) / total
	won /= total
	return fmt.Sprintf("%s: winning ratio: %f losing ratio: %f", p.name, won, lost)
}
