package game

import (
	"fmt"
)

const (
	DRAW = iota - 1
	CONT
	PLAYER1
	PLAYER2
)

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

type Board [9]int

type Game struct {
	board   Board
	player  int
	aiFlags []bool
}

func NewGame(ai1, ai2 bool) *Game {
	return &Game{
		player:  PLAYER1,
		aiFlags: []bool{false, ai1, ai2},
	}
}

func (g *Game) AdvanceTurn() {
	if g.player == PLAYER1 {
		g.player = PLAYER2
	} else {
		g.player = PLAYER1
	}
}

func (g *Game) Render() {
	for i := 0; i < 3; i++ {
		row := g.board[3*i : 3*(i+1)]
		fmt.Printf("%v\n", row)
	}
}

func (g *Game) Move() {
	var i int
	if g.aiFlags[g.player] {
		i, _ = g.AIMove(0)
	} else {
		i = g.PlayerMove()
	}
	g.board[i] = g.player
}

func (g *Game) PlayerMove() int {
	var input string
	var tile int
	fmt.Printf("Player %d to move (1-9): ", g.player)
	fmt.Scanln(&input)
	_, err := fmt.Sscanf(input, "%d", &tile)
	if err != nil || !(1 <= tile && tile <= 9) || g.board[tile-1] != 0 {
		println("Bad input!")
		return g.PlayerMove()
	}
	return tile - 1
}

// Minimax! Player 1 is maximizing.
func (g *Game) AIMove(depth int) (tile, score int) {
	switch g.State() {
	case PLAYER1:
		return -1, 10 - depth
	case PLAYER2:
		return -1, -(10 - depth)
	case DRAW:
		return -1, 0
	}
	var bestTile, bestScore int
	if g.player == PLAYER1 {
		bestScore = minInt
	} else {
		bestScore = maxInt
	}
	for i, tile := range g.board {
		if tile == 0 {
			g.board[i] = g.player
			g.AdvanceTurn()
			_, m := g.AIMove(depth + 1)
			g.AdvanceTurn()
			g.board[i] = 0
			if (g.player == PLAYER1 && bestScore < m) || (g.player == PLAYER2 && m < bestScore) {
				bestTile, bestScore = i, m
			}
		}
	}
	return bestTile, bestScore
}

func (g *Game) State() int {
	tiles := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, //Diags
	}
	for _, t := range tiles {
		a, b, c := g.board[t[0]], g.board[t[1]], g.board[t[2]]
		for _, p := range []int{PLAYER1, PLAYER2} {
			if p == a && a == b && b == c {
				return p
			}
		}
	}
	for _, i := range g.board {
		if i == 0 {
			return CONT
		}
	}
	return DRAW
}
