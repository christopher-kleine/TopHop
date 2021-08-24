package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/tilesets/world.png
var worldPng []byte

type World struct {
	player  *Player
	level   *TMX
	tileset *ebiten.Image
	Game    *Game
	startX  []int
	startY  []int
	goalX   []int
	goalY   []int
}

func (w *World) Update() error {
	w.player.Update()

	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	//screen.DrawImage(w.tileset, nil)
	w.player.Draw(screen)
}

func NewWorld(g *Game) *World {
	var err error
	worldDecoded, _, err := image.Decode(bytes.NewReader(worldPng))
	if err != nil {
		log.Fatal(err)
	}

	return &World{
		Game:    g,
		tileset: ebiten.NewImageFromImage(worldDecoded),
		level:   &TMX{},
		startX:  []int{1, 1, 1, 1, 1, 1, 1},
		startY:  []int{7, 7, 7, 7, 12, 12, 9, 2},
		goalX:   []int{29, 29, 29, 29, 29, 29, 29, 29},
		goalY:   []int{7, 7, 7, 12, 12, 9, 2, 4},
		player:  NewPlayer(func(x, y int) bool { return (x < 10) || (x >= 10 && y >= 15) }),
	}
}
