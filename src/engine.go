package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var tickCount = 0

func debugDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.0f", ebiten.ActualFPS()))

	// Draw game map
	tileSize := scrHeight / gMapGridHeight
	xPadding := 55 // x-axis padding to account for wide window
	for i := 0; i < gMapGridHeight; i++ {
		for j := 0; j < gMapGridWidth; j++ {
			if gMap[i][j] == 1 {
				ebitenutil.DrawRect(screen, float64(xPadding+(j*tileSize)), float64(i*tileSize), float64(tileSize-1), float64(tileSize-1), color.White) // square size so that tiles appear in a grid
			}
		}
	}

	// Draw player
	ebitenutil.DrawCircle(screen, player.xPos, player.yPos, 3, color.RGBA{255, 0, 0, 255})

	// Draw line of sight
	//degr := (player.angle * 180) / math.Pi
	fmt.Println(player.angle)
	fmt.Println(math.Cos(player.angle))
	ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(player.angle)*20), player.yPos+(math.Sin(player.angle)*20), color.RGBA{0, 0, 255, 255})
}

// TODO refactor to factor in dx,dy movement
func collision(xPos, yPos int) bool {
	// Get 'tile' player is on
	x := ((xPos - 55) / (scrHeight / gMapGridHeight))
	y := ((yPos) / (scrHeight / gMapGridHeight))
	fmt.Printf("On tile [%d][%d]", x, y)
	fmt.Println()
	return gMap[y][x] == 1
}
