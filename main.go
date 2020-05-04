package main

import (
	"fmt"
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
	g.Render()
	if state == game.DRAW {
		println("A shameful display!")
	} else {
		fmt.Printf("Player %d won!", state)
	}
}
