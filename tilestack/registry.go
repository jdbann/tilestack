package tilestack

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Registry struct {
	tiles    []TileDefinition
	tileSize float32
}

func NewRegistry() *Registry {
	return &Registry{
		tileSize: 16,
	}
}

func (r *Registry) Load(name string, steps int32) int {
	fileName := fmt.Sprintf("assets/%s_strip%d.png", name, steps)
	texture := rl.LoadTexture(fileName)
	r.tiles = append(r.tiles, TileDefinition{
		texture: texture,
		frames:  steps,
		size: rl.Vector2{
			X: float32(texture.Width / steps),
			Y: float32(texture.Height),
		},
	})
	return len(r.tiles) - 1
}

func (r *Registry) DrawMap(tiles TileMap, angle float32) {
	x, y, z := tiles.Size()
	origin := rl.Vector3{X: float32(x-1) / -2 * r.tileSize, Y: float32(y-1) / -2 * r.tileSize, Z: float32(z-1) / -2 * r.tileSize}

	worldTranslation := rl.MatrixMultiply(
		rl.MatrixScale(r.tileSize, r.tileSize, r.tileSize),
		rl.MatrixMultiply(
			rl.MatrixTranslate(origin.X, origin.Y, 0),
			rl.MatrixRotateZ(-angle),
		),
	)

	yFrom, yTo := 0, len(tiles[0])-1
	if math.Cos(float64(angle)) < 0 {
		yFrom, yTo = yTo, yFrom
	}

	xFrom, xTo := 0, len(tiles[0][0])-1
	if math.Sin(float64(angle)) < 0 {
		xFrom, xTo = xTo, xFrom
	}

	yScale := float32(math.Cos(rl.Pi / 3))
	rl.PushMatrix()
	defer rl.PopMatrix()
	rl.Scalef(1, yScale, 1)

	for z := 0; z < len(tiles); z++ {
		for yNext, yDone := iterator(yFrom, yTo); !yDone(); {
			y := yNext()
			for xNext, xDone := iterator(xFrom, xTo); !xDone(); {
				x := xNext()
				p := rl.Vector3Transform(rl.Vector3{X: float32(x), Y: float32(y), Z: float32(z)}, worldTranslation)
				r.DrawTile(tiles[z][y][x], p, angle, 1/yScale)
			}
		}
	}
}

func (r *Registry) DrawTile(t Tile, p rl.Vector3, angle, step float32) {
	if t.index == -1 {
		return
	}

	tile := r.tiles[t.index]
	origin := rl.NewVector2(tile.size.X/2, tile.size.Y/2)

	for frame := float32(0); frame < float32(tile.frames); frame++ {
		rl.DrawTexturePro(
			tile.texture,
			rl.NewRectangle(frame*tile.size.X, 0, tile.size.X, tile.size.Y),
			rl.NewRectangle(p.X, p.Y-(p.Z*step)-(frame*step), tile.size.X, tile.size.Y),
			origin,
			(angle+float32(t.dir))*rl.Rad2deg,
			rl.White,
		)
	}
}

type TileDefinition struct {
	texture rl.Texture2D
	frames  int32
	size    rl.Vector2
}

func iterator(from, to int) (func() int, func() bool) {
	v, step := from, 1
	if from > to {
		step = -1
	}
	done := false

	nextFn := func() int {
		out := v
		if out == to {
			done = true
		}
		v += step
		return out
	}

	doneFn := func() bool { return done }

	return nextFn, doneFn
}

type TileDirection float32

const (
	North TileDirection = 0
	East                = rl.Pi / 2
	South               = rl.Pi
	West                = 3 * rl.Pi / 2
)

type Tile struct {
	index int
	dir   TileDirection
}

type TileMap [][][]Tile

func NewTileMap(x, y, z int) TileMap {
	m := make(TileMap, z)
	for cz := 0; cz < z; cz++ {
		m[cz] = make([][]Tile, y)
		for cy := 0; cy < y; cy++ {
			m[cz][cy] = make([]Tile, x)
			for cx := 0; cx < x; cx++ {
				m[cz][cy][cx] = Tile{-1, 0}
			}
		}
	}
	return m
}

func (m TileMap) Set(x, y, z, idx int, dir TileDirection) {
	m[z][y][x].index = idx
	m[z][y][x].dir = dir
}

func (m TileMap) At(x, y, z int) Tile {
	return m[z][y][x]
}

func (m TileMap) Rect(x0, x1, y0, y1, z, idx int) {
	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			m.Set(x, y, z, idx, 0)
		}
	}
}

func (m TileMap) Size() (int, int, int) {
	return len(m[0][0]), len(m[0]), len(m)
}
