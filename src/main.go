package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var tick = 0

func (g *Game) Update() error {
	movePlayer()
	tick++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	debugDraw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scrWidth, scrHeight
}

func main() {
	loadImages()
	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowTitle("Raycasting - Go")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
