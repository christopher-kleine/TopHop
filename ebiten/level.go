package main

import (
	"bytes"
	"image"
	"log"
	"net/http"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type Level struct {
	level   *tiled.Map
	loader  *tiled.Loader
	tileset *ebiten.Image
	PlayerX int
	PlayerY int
	GoalX   int
	GoalY   int
}

func (l *Level) IsSolid(X, Y int) bool {
	if Y > 19 || X > 29 {
		return false
	}

	id := Y*30 + X
	tile := l.level.Layers[1].Tiles[id]

	return !tile.Nil && !l.IsGoal(X, Y)
}

func (l *Level) IsGoal(X, Y int) bool {
	if Y > 19 || X > 29 {
		return false
	}

	id := Y*30 + X
	tile := l.level.Layers[1].Tiles[id].ID

	return tile == 4
}

func (l *Level) IsSpring(X, Y int) bool {
	if Y > 19 || X > 29 {
		return false
	}

	id := Y*30 + X
	tile := l.level.Layers[1].Tiles[id].ID

	return tile == 9
}

func (l *Level) IsSpringable(X, Y int) bool {
	if Y > 19 || X > 29 {
		return false
	}

	id := Y*30 + X
	tile := l.level.Layers[1].Tiles[id].ID

	return tile == 5
}

func (l *Level) IsDeadly(X, Y int) bool {
	if Y > 19 || X > 29 {
		return true
	}

	id := Y*30 + X
	tile := l.level.Layers[1].Tiles[id].ID

	return tile == 7
}

func (l *Level) Draw(screen *ebiten.Image) {
	var (
		x int
		y int
	)
	for _, tile := range l.level.Layers[0].Tiles {
		if tile.Nil == false {
			spriteRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := l.tileset.SubImage(spriteRect).(*ebiten.Image)

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x)*8, float64(y)*8)
			screen.DrawImage(tileImage, op)
		}
		x = (x + 1) % 30
		if x == 0 {
			y++
		}
	}
}

func (l *Level) Load(filename string) error {
	var err error

	l.level, err = l.loader.LoadFromFile("assets/maps/" + filename)
	if err != nil {
		return err
	}

	var (
		x int
		y int
	)
	for _, tile := range l.level.Layers[1].Tiles {
		if tile.Nil == false {
			switch tile.ID {
			case 3:
				l.PlayerX = x
				l.PlayerY = y

			case 4:
				l.GoalX = x
				l.GoalY = y
			}
		}
		x = (x + 1) % 30
		if x == 0 {
			y++
		}
	}

	return nil
}

func NewLevel(filename string) *Level {
	var err error
	worldPng, err := assets.ReadFile("assets/images/world.png")
	if err != nil {
		log.Fatal(err)
	}

	worldDecoded, _, err := image.Decode(bytes.NewReader(worldPng))
	if err != nil {
		log.Fatal(err)
	}

	level := &Level{
		tileset: ebiten.NewImageFromImage(worldDecoded),
		loader:  &tiled.Loader{FileSystem: http.FS(assets)},
	}

	if err := level.Load(filename); err != nil {
		log.Fatal(err)
	}

	return level
}
