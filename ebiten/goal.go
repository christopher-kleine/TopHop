package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Goal struct {
	game *Game
	goal *Level
}

func (g *Goal) Reset() {}

func (g *Goal) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.game.SetScene(SplashScene, true)
	}

	return nil
}

func (g *Goal) Draw(screen *ebiten.Image) {
	g.goal.Draw(screen)
}

func NewGoal(game *Game) *Goal {
	return &Goal{
		game: game,
		goal: NewLevel("goal.tmx"),
	}
}
