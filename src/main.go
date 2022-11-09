package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	winWidth       = 1600
	winHeight      = 900
	scrWidth       = winWidth / 2
	scrHeight      = winHeight / 2
	gMapGridWidth  = 15
	gMapGridHeight = 10
)

// Game map
// 1 - Wall
// 2 - Empty space
var gMap = [gMapGridHeight][gMapGridWidth]int{
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{1, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.0f", ebiten.ActualFPS()))

	// Draw game map
	squareSize := scrHeight / gMapGridHeight
	xPadding := 55 // x-axis padding to account for wide window
	for i := 0; i < gMapGridHeight; i++ {
		for j := 0; j < gMapGridWidth; j++ {
			if gMap[i][j] == 1 {
				ebitenutil.DrawRect(screen, float64(xPadding+(j*squareSize)), float64(i*squareSize), float64(squareSize-1), float64(squareSize-1), color.White)
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return scrWidth, scrHeight
}

func main() {

	ebiten.SetWindowSize(winWidth, winHeight)
	ebiten.SetWindowTitle("Raycasting - Go")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
