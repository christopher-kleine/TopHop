package main

import "github.com/hajimehoshi/ebiten/v2"

type Pause struct {
	game *Game
}

func (p *Pause) Update() error {
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
