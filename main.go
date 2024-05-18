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
	floorPlain := reg.Load("FloorCentrePlain", 16)
	floorEdge := reg.Load("FloorEdgeSide", 16)
	floorCorner := reg.Load("FloorEdgeCorner", 16)
	floorPattern := reg.Load("FloorCentrePattern", 16)
	floorPatternEdge := reg.Load("FloorPatternFade", 16)
	floorPatternCorner := reg.Load("FloorPatternFadeCornerInt", 16)
	wallCorner := reg.Load("WallCorner", 16)
	wallCornerInt := reg.Load("WallCornerInt", 16)
	wallSide := reg.Load("WallSideE", 16)
	door := reg.Load("TwoSidedDoor", 16)
	table := reg.Load("Table", 8)
	chair := reg.Load("Chair", 12)

	tileMap := tilestack.NewTileMap(7, 9, 3)

	// Floor
	tileMap.RectWithEdges(0, 6, 0, 8, 0, floorPlain, floorEdge, floorCorner)
	tileMap.RectWithEdges(1, 5, 1, 4, 0, floorPattern, floorPatternEdge, floorPatternCorner)

	// Walls
	tileMap.Set(0, 8, 1, wallCornerInt, tilestack.West)
	tileMap.Set(1, 8, 1, wallSide, tilestack.East)
	tileMap.Set(2, 8, 1, wallCorner, tilestack.South)
	tileMap.Set(4, 8, 1, wallCorner, tilestack.East)
	tileMap.Set(5, 8, 1, wallSide, tilestack.East)
	tileMap.Set(6, 8, 1, wallCornerInt, tilestack.North)
	tileMap.Set(6, 7, 1, wallSide, tilestack.North)
	tileMap.Set(6, 6, 1, wallSide, tilestack.North)
	tileMap.Set(6, 5, 1, wallSide, tilestack.North)

	// Objects
	tileMap.Set(3, 8, 1, door, tilestack.North)
	tileMap.Set(2, 2, 1, table, tilestack.North)
	tileMap.Set(2, 3, 1, chair, tilestack.North)
	tileMap.Set(3, 2, 1, chair, tilestack.West)

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
