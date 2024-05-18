package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdbann/tilestack/tilestack"
)

const (
	screenWidth, screenHeight   int32 = 1280, 720
	virtualWidth, virtualHeight int32 = screenWidth / 2, screenHeight / 2
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "tilestack")
	defer rl.CloseWindow()

	reg := tilestack.NewRegistry()
	reg.Load("FloorCentrePattern", 16)
	reg.Load("Table", 8)
	reg.Load("Chair", 12)

	tileMap := tilestack.NewMap(7, 9, 3)
	tileMap.Rect(0, 6, 0, 8, 0, 0)
	tileMap.Set(1, 1, 1, 1)
	tileMap.Set(1, 2, 1, 2)

	virtualScreen := rl.LoadRenderTexture(virtualWidth, virtualHeight)
	virtualScreenRec := rl.NewRectangle(0, 0, float32(virtualWidth), -float32(virtualHeight))
	screenRec := rl.NewRectangle(0, 0, float32(screenWidth), float32(screenHeight))

	for !rl.WindowShouldClose() {
		cameraAngle := rl.GetMousePosition().X * rl.Pi * 2 / float32(screenWidth)

		rl.BeginTextureMode(virtualScreen)
		rl.ClearBackground(rl.RayWhite)
		rl.Translatef(float32(virtualWidth)/2, float32(virtualHeight)/2, 0)
		reg.DrawMap(tileMap, cameraAngle)
		rl.EndTextureMode()

		rl.BeginDrawing()
		rl.DrawTexturePro(virtualScreen.Texture, virtualScreenRec, screenRec, rl.Vector2{}, 0, rl.White)
		rl.EndDrawing()
	}
}
