package main

import (
	"net"

	"github.com/ARF-DEV/ping-pong-mp/common/core"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(1000, 500, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	data := []byte("halloww")
	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}

	rl.SetTargetFPS(60)
	game := core.CreateGame()
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		game.Update()
		game.Draw()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}
}
