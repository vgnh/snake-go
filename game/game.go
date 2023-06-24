package game

import "time"

const (
	TEN = 10
)

type direction int32

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

type point struct {
	X, Y int32
}

type snake struct {
	Body      []point
	direction direction
}

type Game struct {
	Canvas [][]int
	Snake  snake
	Mochi  point
	speed  int
}

func NewGame(width, height int) *Game {
	canvas := make([][]int, height/TEN)
	for i := range canvas {
		canvas[i] = make([]int, width/TEN)
	}
	snake := snake{
		[]point{{20, 0}, {21, 0}},
		WEST,
	}
	mochi := point{10, 0}

	return &Game{
		canvas,
		snake,
		mochi,
		500,
	}
}

func (g *Game) UpdateGameState() {
	for i := range g.Snake.Body {
		switch g.Snake.direction {
		case NORTH:
			g.Snake.Body[i].Y -= 1
		case EAST:
			g.Snake.Body[i].X += 1
		case SOUTH:
			g.Snake.Body[i].Y += 1
		case WEST:
			g.Snake.Body[i].X -= 1
		}
	}

	time.Sleep(time.Millisecond * time.Duration(g.speed))
}
