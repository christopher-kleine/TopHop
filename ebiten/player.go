package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/player.png
var playerPng []byte

type Player struct {
	tileset *ebiten.Image
}

func (p *Player) Draw(screen *ebiten.Image) {
	screen.DrawImage(p.tileset, nil)
}

func NewPlayer() *Player {
	var err error
	playerDecoded, _, err := image.Decode(bytes.NewReader(playerPng))
	if err != nil {
		log.Fatal(err)
	}

	return &Player{
		tileset: ebiten.NewImageFromImage(playerDecoded),
	}
}
