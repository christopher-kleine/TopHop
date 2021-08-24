package main

import (
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type World struct {
	player  *Player
	level   *Level
	game    *Game
	levelID int
	startX  []int
	startY  []int
	goalX   []int
	goalY   []int
}

func (w *World) Reset() {
	w.levelID = 0
	err := w.level.Load(fmt.Sprintf("level%d.tmx", w.levelID))
	if err != nil {
		log.Fatal(err)
	}

	w.player.PosX = w.level.PlayerX
	w.player.PosY = w.level.PlayerY
}

func (w *World) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		w.game.SetScene(PauseScene, false)
		return nil
	}

	w.player.Update()

	if w.player.PosX == w.level.GoalX && w.player.PosY == w.level.GoalY {
		w.levelID++
		err := w.level.Load(fmt.Sprintf("level%d.tmx", w.levelID))
		if err != nil {
			w.game.SetScene(GoalScene, true)
		} else {
			w.player.PosX = w.level.PlayerX
			w.player.PosY = w.level.PlayerY
		}
	}

	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	w.level.Draw(screen)
	w.player.Draw(screen)
}

func NewWorld(g *Game) *World {
	// isSolidDummy := func(x, y int) bool {
	// 	return (x < 9) || (x >= 9 && y >= 18)
	// }

	l := NewLevel("level1.tmx")

	world := &World{
		game:    g,
		startX:  []int{1, 1, 1, 1, 1, 1, 1},
		startY:  []int{7, 7, 7, 7, 12, 12, 9, 2},
		goalX:   []int{29, 29, 29, 29, 29, 29, 29, 29},
		goalY:   []int{7, 7, 7, 12, 12, 9, 2, 4},
		player:  NewPlayer(l.IsSolid),
		level:   l,
		levelID: 0,
	}

	return world
}
