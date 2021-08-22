package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Splash struct {
	Game *Game
}

func (s *Splash) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.Game.Scene = "levels"
	}

	return nil
}

func (s *Splash) Draw(screen *ebiten.Image) {
	s.Game.World.Draw(screen)
}
