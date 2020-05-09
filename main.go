package main

import "tictacgo/game"

func main() {
	g := game.New(false, true)
	for g.State() == game.Cont {
		g.Move()
		g.Render()
		g.Advance()
	}
	if state := g.State(); state == game.Draw {
		println("A shameful display!")
	} else {
		println("Player", state, "won!")
	}
}
