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
	scenes  map[SceneID]Scene
	sceneID SceneID
}

func (g *Game) SetScene(sceneID SceneID, reset bool) {
	if reset {
		g.scenes[sceneID].Reset()
	}
	g.sceneID = sceneID
}

func (g *Game) Update() error {
	return g.scenes[g.sceneID].Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scenes[g.sceneID].Draw(screen)

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
		scenes:  make(map[SceneID]Scene),
		sceneID: SplashScene,
	}

	game.scenes[PauseScene] = NewPause(game)
	game.scenes[SplashScene] = NewSplash(game)
	game.scenes[WorldScene] = NewWorld(game)
	game.scenes[GoalScene] = NewGoal(game)

	game.scenes[WorldScene].Reset()

	return game
}
