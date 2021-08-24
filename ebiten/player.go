package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/tilesets/player.png
var playerPng []byte

type Player struct {
	tileset *ebiten.Image
	IsSolid func(X, Y int) bool
	State   func() error

	frameCount int
	animFrame  int
	Speed      int

	PosX    int
	PosY    int
	OffsetX float64
	OffsetY float64
	velX    float64
	velY    float64
}

func (p *Player) Update() error {
	return p.State()
}

func (p *Player) Draw(screen *ebiten.Image) {
	tileX := p.animFrame % 4
	tileY := p.animFrame / 4
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.PosX*8)+p.OffsetX, float64(p.PosY*8)+p.OffsetY)
	tile := image.Rect(tileX*8, tileY*8, tileX*8+8, tileY*8+8)
	screen.DrawImage(p.tileset.SubImage(tile).(*ebiten.Image), op)
}

func (p *Player) Walk() error {
	p.frameCount = (p.frameCount + 1) % p.Speed
	if p.frameCount == 0 {
		p.animFrame = (p.animFrame + 1) % 4
	}
	p.OffsetX = p.OffsetX + p.velX
	if p.OffsetX >= 8 {
		tileCount := int(p.OffsetX) / 8
		p.PosX = p.PosX + tileCount
		p.OffsetX = p.OffsetX - float64(tileCount)*8
	}
	if !p.IsSolid(p.PosX, p.PosY) {
		p.State = p.Fall
		p.frameCount = 0
		p.animFrame = 4
	}

	return nil
}

func (p *Player) Jump() error {
	p.frameCount = (p.frameCount + 1) % 5
	if p.frameCount == 0 {
		p.animFrame = (p.animFrame+1)%2 + 6
	}
	// p.OffsetX = p.OffsetX + p.velX
	// if p.OffsetX >= 8 {
	// 	tileCount := int(p.OffsetX) / 8
	// 	p.PosX = p.PosX + tileCount
	// 	p.OffsetX = p.OffsetX - float64(tileCount)*8
	// }
	// if !p.IsSolid(p.PosX, p.PosY) {
	// 	p.State = p.Fall
	// }

	return nil
}

func (p *Player) Fall() error {
	p.frameCount = (p.frameCount + 1) % 5
	if p.frameCount == 0 {
		p.animFrame = (p.animFrame+1)%2 + 4
	}
	p.OffsetY = p.OffsetY + p.velY
	if p.OffsetY >= 8 {
		tileCount := int(p.OffsetY) / 8
		p.PosY = p.PosY + tileCount
		p.OffsetY = p.OffsetY - float64(tileCount)*8
	}
	if p.IsSolid(p.PosX, p.PosY) {
		p.State = p.Walk
		p.frameCount = 0
		p.animFrame = 0
	}

	return nil
}

func NewPlayer(IsSolid func(int, int) bool) *Player {
	var err error
	playerDecoded, _, err := image.Decode(bytes.NewReader(playerPng))
	if err != nil {
		log.Fatal(err)
	}

	p := &Player{
		tileset: ebiten.NewImageFromImage(playerDecoded),
		Speed:   15,
		velX:    0.5,
		velY:    1,
		IsSolid: IsSolid,
	}

	p.State = p.Walk

	return p
}
