package main

import (
	"tictacgo/game"
)

func main() {
	g := game.NewGame(false, true)
	var state int
	for state = game.CONT; state == game.CONT; state = g.State() {
		g.Move()
		g.Render()
		g.AdvanceTurn()
	}
	if state == game.DRAW {
		println("A shameful display!")
	} else {
		println("Player", state, "won!")
	}
}
