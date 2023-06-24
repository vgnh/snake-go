package main

import (
	"snake-go/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH = 800
	HEIGHT = 450
)

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "snake go")

	rl.SetTargetFPS(60)

	game := game.NewGame(WIDTH, HEIGHT)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.DrawText("This is the base canvas!", WIDTH/3, HEIGHT/2.25, 20, rl.White)

		drawGameUI(game)
		game.UpdateGameState()

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func drawGameUI(game *game.Game) {
	// Draw canvas
	for i := range game.Canvas {
		for j := range game.Canvas[i] {
			rl.DrawRectangle(int32(j)*10, int32(i)*10, 10, 10, rl.RayWhite)
		}
	}

	// Draw snake
	for _, point := range game.Snake.Body {
		rl.DrawRectangle(point.X*10, point.Y*10, 10, 10, rl.SkyBlue)
	}

	// Draw mochi
	rl.DrawRectangle(game.Mochi.X*10, game.Mochi.Y*10, 10, 10, rl.Orange)
}
