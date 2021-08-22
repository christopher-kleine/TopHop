package main

import "github.com/hajimehoshi/ebiten/v2"

type Pause struct {
	Game *Game
}

func (p *Pause) Update() error {
	return nil
}

func (p *Pause) Draw(screen *ebiten.Image) {
	p.Game.Player.Draw(screen)
}
