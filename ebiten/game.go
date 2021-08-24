package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneID int

const (
	SplashScene SceneID = iota
	PauseScene
	WorldScene
	GoalScene
)

type Game struct {
	scenes map[SceneID]Scene
	Scene  SceneID
}

func (g *Game) Update() error {
	return g.scenes[g.Scene].Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scenes[g.Scene].Draw(screen)

	//ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %.0f\nTPS: %.0f", ebiten.CurrentFPS(), ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideW, outsideH int) (int, int) {
	return 240, 160
}

func (g *Game) Render(scene SceneID, screen *ebiten.Image) {
	g.scenes[scene].Draw(screen)
}

func NewGame() *Game {
	game := &Game{
		scenes: make(map[SceneID]Scene),
		Scene:  SplashScene,
	}

	game.scenes[PauseScene] = NewPause(game)
	game.scenes[SplashScene] = NewSplash(game)
	game.scenes[WorldScene] = NewWorld(game)

	return game
}
