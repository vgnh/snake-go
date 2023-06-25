package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/vgnh/snake-go/game"
)

const (
	WIDTH  = 800
	HEIGHT = 450
)

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "snake go")
	rl.SetTargetFPS(12)

	game := game.NewGame(WIDTH, HEIGHT)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.DrawText("This is the base canvas!", WIDTH/3, HEIGHT/2.25, 20, rl.White)

		switch rl.GetKeyPressed() {
		case rl.KeyUp:
			game.UpdateSnakeDirection(game.Up())
		case rl.KeyRight:
			game.UpdateSnakeDirection(game.Right())
		case rl.KeyDown:
			game.UpdateSnakeDirection(game.Down())
		case rl.KeyLeft:
			game.UpdateSnakeDirection(game.Left())
		case rl.KeyEscape:
			break
		case 0:
		}
		game.UpdateGameState()
		drawGameUI(game)

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

	// Draw score
	rl.DrawText(fmt.Sprintf("Score: %v", game.Score), 0, 0, 20, rl.Beige)

	rl.DrawText("Press ESC to exit", WIDTH/2.5, 0, 20, rl.LightGray)

	// Draw snake
	for _, point := range game.Snake.Body {
		rl.DrawRectangle(int32(point.X*10), int32(point.Y*10), 10, 10, rl.SkyBlue)
	}

	// Draw mochi
	rl.DrawCircle(int32((game.Mochi.X*10)+5), int32((game.Mochi.Y*10)+5), 5, rl.Orange)
}
