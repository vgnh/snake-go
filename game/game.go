package game

import (
	"math/rand"
)

const (
	TEN = 10
)

type direction int32

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

type point struct {
	X, Y int
}

type snake struct {
	Body      []point
	direction direction
}

type Game struct {
	Canvas [][]int
	Snake  snake
	Mochi  point
	Score  int
}

func NewGame(width, height int) *Game {
	canvas := make([][]int, height/TEN)
	for i := range canvas {
		canvas[i] = make([]int, width/TEN)
	}
	snake := snake{
		[]point{{20, 20}, {21, 20}, {22, 20}},
		LEFT,
	}
	mochi := point{10, 20}

	return &Game{
		canvas,
		snake,
		mochi,
		0,
	}
}

func (g *Game) UpdateGameState() {
	newPoint := g.Snake.Body[0]
	switch g.Snake.direction {
	case UP:
		newPoint.Y -= 1
	case RIGHT:
		newPoint.X += 1
	case DOWN:
		newPoint.Y += 1
	case LEFT:
		newPoint.X -= 1
	}

	if (newPoint.X >= len(g.Canvas[0]) || newPoint.X < 0) || (newPoint.Y >= len(g.Canvas) || newPoint.Y < 0) {
		return
	}

	if newPoint == g.Mochi {
		g.Snake.Body = append([]point{g.Mochi}, g.Snake.Body[0:]...)
		g.Mochi = g.newMochiPosition()
		g.Score += 1
	} else {
		g.Snake.Body = append([]point{newPoint}, g.Snake.Body[:len(g.Snake.Body)-1]...)
	}
}

func (g *Game) newMochiPosition() point {
	for {
		y := rand.Intn(len(g.Canvas))
		x := rand.Intn(len(g.Canvas[y]))

		pointIsSafe := true
		for _, v := range g.Snake.Body {
			if v.X == x && v.Y == y {
				pointIsSafe = false
				break
			}
		}

		if pointIsSafe {
			return point{
				x,
				y,
			}
		}
	}
}

func (g *Game) UpdateSnakeDirection(d direction) {
	switch {
	case g.Snake.direction == UP && d == DOWN:
		fallthrough
	case g.Snake.direction == DOWN && d == UP:
		fallthrough
	case g.Snake.direction == LEFT && d == RIGHT:
		fallthrough
	case g.Snake.direction == RIGHT && d == LEFT:
		fallthrough
	case g.Snake.direction == d:
		return
	default:
		g.Snake.direction = d
	}
}

func (g *Game) Up() direction {
	return UP
}

func (g *Game) Down() direction {
	return DOWN
}

func (g *Game) Left() direction {
	return LEFT
}

func (g *Game) Right() direction {
	return RIGHT
}
