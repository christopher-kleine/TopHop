package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Pause struct {
	game *Game
}

func (p *Pause) Reset() {}

func (p *Pause) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		p.game.SetScene(WorldScene, false)
	}

	if ids := inpututil.JustPressedTouchIDs(); ids != nil {
		p.game.SetScene(WorldScene, false)
	}

	return nil
}

func (p *Pause) Draw(screen *ebiten.Image) {
	p.game.Render(WorldScene, screen)
}

func NewPause(game *Game) *Pause {
	return &Pause{
		game: game,
	}
}
