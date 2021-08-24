package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Splash struct {
	game   *Game
	splash *Level
}

func (s *Splash) Reset() {}

func (s *Splash) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.game.SetScene(PauseScene, true)
	}

	return nil
}

func (s *Splash) Draw(screen *ebiten.Image) {
	s.splash.Draw(screen)
}

func NewSplash(game *Game) *Splash {
	return &Splash{
		game:   game,
		splash: NewLevel("splash.tmx"),
	}
}
