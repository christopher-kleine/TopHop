package main

import (
	"embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// go:embed assets/* assets/maps/* assets/tilesets/*
var assets embed.FS

func main() {
	game := NewGame()
	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(240*4, 160*4)
	ebiten.SetWindowTitle("Top Hop - A tiny challenge")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
