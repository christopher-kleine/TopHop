package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/world.png
var worldPng []byte

type World struct {
	tileset *ebiten.Image
}

func (w *World) Draw(screen *ebiten.Image) {
	screen.DrawImage(w.tileset, nil)
}

func NewWorld() *World {
	var err error
	worldDecoded, _, err := image.Decode(bytes.NewReader(worldPng))
	if err != nil {
		log.Fatal(err)
	}

	return &World{
		tileset: ebiten.NewImageFromImage(worldDecoded),
	}
}
