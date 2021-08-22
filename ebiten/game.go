package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	Player *Player
	World  *World

	scenes map[string]Scene
	Scene  string
}

func (g *Game) SetScene(name string) {
}

func (g *Game) Update() error {
	return g.scenes[g.Scene].Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scenes[g.Scene].Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.0f\nTPS: %.0f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
	return 240, 160
}

func NewGame() *Game {
	game := &Game{
		Player: NewPlayer(),
		World:  NewWorld(),
		scenes: make(map[string]Scene),
		Scene:  "splash",
	}

	game.scenes["pause"] = &Pause{Game: game}
	game.scenes["splash"] = &Splash{Game: game}
	game.scenes["levels"] = &Levels{Game: game}

	return game
}
