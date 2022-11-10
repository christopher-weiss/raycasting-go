package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func draw(screen *ebiten.Image) {
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

	// Draw player
	ebitenutil.DrawRect(screen, player.xPos, player.yPos, 3, 3, color.RGBA{255, 0, 0, 255})
}

func collision(xPos, yPos int) bool {
	// Get 'tile' player is on
	x := ((xPos - 55) / (scrHeight / gMapGridHeight))
	y := ((yPos) / (scrHeight / gMapGridHeight))
	fmt.Printf("On tile [%d][%d]", x, y)
	fmt.Println()
	return gMap[y][x] == 1
}
