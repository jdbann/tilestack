package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/tilestack/tilestack"
)

const (
	screenWidth, screenHeight int32 = 1280, 720
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "tilestack")
	defer rl.CloseWindow()

	reg := tilestack.NewRegistry()
	reg.Load("FloorCentrePlain", 16)
	reg.Load("Table", 8)
	reg.Load("Chair", 12)

	tileMap := tilestack.NewMap(3, 3, 3)
	tileMap.Rect(0, 2, 0, 2, 0, 0)
	tileMap.Set(1, 1, 1, 1)
	tileMap.Set(1, 2, 1, 2)

	for !rl.WindowShouldClose() {
		cameraAngle := rl.GetMousePosition().X * rl.Pi * 2 / float32(screenWidth)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.Translatef(float32(screenWidth)/2, float32(screenHeight)/2, 0)

		reg.DrawMap(tileMap, cameraAngle)

		rl.EndDrawing()
	}
}
