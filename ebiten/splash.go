package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Splash struct {
	game *Game
}

func (s *Splash) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.game.Scene = WorldScene
	}

	return nil
}

func (s *Splash) Draw(screen *ebiten.Image) {
}

func NewSplash(game *Game) *Splash {
	return &Splash{
		game: game,
	}
}
