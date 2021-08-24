package main

import "github.com/hajimehoshi/ebiten/v2"

type Scene interface {
	Reset()
	Update() error
	Draw(*ebiten.Image)
}
