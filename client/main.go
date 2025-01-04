package main

import (
	"github.com/ARF-DEV/ping-pong-mp/common/core"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	game := core.CreateGame("client")
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		game.Update()
		game.Draw()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}
}
