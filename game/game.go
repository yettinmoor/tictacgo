package game

import "fmt"

type state int

const (
	Draw state = iota - 1
	Cont
	Player1
	Player2
)

type Game struct {
	board   [9]state
	player  state
	aiFlags []bool
}

func New(ai1, ai2 bool) *Game {
	return &Game{
		player:  Player1,
		aiFlags: []bool{false, ai1, ai2},
	}
}

func (g *Game) Advance() {
	g.player = map[state]state{
		Player1: Player2,
		Player2: Player1,
	}[g.player]
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
		i, _ = g.aiMove(0)
	} else {
		i = g.playerMove()
	}
	g.board[i] = g.player
}

func (g *Game) playerMove() int {
	var input string
	var tile int
	fmt.Printf("Player %d to move (1-9): ", g.player)
	fmt.Scanln(&input)
	_, err := fmt.Sscanf(input, "%d", &tile)
	if err != nil || !(1 <= tile && tile <= 9) || g.board[tile-1] != 0 {
		println("Bad input!")
		return g.playerMove()
	}
	return tile - 1
}

// Minimax! Player 1 is maximizing.
func (g *Game) aiMove(depth int) (optTile, score int) {
	switch g.State() {
	case Player1:
		score = 10 - depth
		return
	case Player2:
		score = depth - 10
		return
	case Draw:
		return
	}
	score = map[state]int{
		Player1: -20,
		Player2: 20,
	}[g.player]
	for i, tile := range g.board {
		if tile == 0 {
			g.board[i] = g.player
			g.Advance()
			_, m := g.aiMove(depth + 1)
			g.Advance()
			g.board[i] = 0
			if (g.player == Player1 && score < m) || (g.player == Player2 && m < score) {
				optTile, score = i, m
			}
		}
	}
	return
}

func (g *Game) State() state {
	tiles := [][3]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // Rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // Columns
		{0, 4, 8}, {2, 4, 6}, // Diags
	}
	for _, t := range tiles {
		a, b, c := g.board[t[0]], g.board[t[1]], g.board[t[2]]
		if 0 != a && a == b && b == c {
			return a
		}
	}
	for _, i := range g.board {
		if i == 0 {
			return Cont
		}
	}
	return Draw
}
