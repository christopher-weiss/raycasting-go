package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func debugDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.0f", ebiten.ActualFPS()))

	// Draw game map
	tileSize := scrHeight / gMapGridHeight
	xPadding := 55 // x-axis padding to account for wide window
	for i := 0; i < gMapGridHeight; i++ {
		for j := 0; j < gMapGridWidth; j++ {
			if gMap[i][j] == 1 {
				ebitenutil.DrawRect(screen, float64(xPadding+(j*tileSize)), float64(i*tileSize), float64(tileSize-1), float64(tileSize-1), color.White)
			}
		}
	}

	// Draw player
	ebitenutil.DrawCircle(screen, player.xPos, player.yPos, 3, color.RGBA{255, 0, 0, 255})

	// Draw line of sight
	ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(player.angle)*20), player.yPos+(math.Sin(player.angle)*20), color.RGBA{0, 0, 255, 255})

	// Draw rays
	rayAngle := player.angle - halfFov + 0.0001
	for i := 0; i < numRays; i++ {
		rayAngle += deltaAngle
		ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(rayAngle)*maxDepth), player.yPos+(math.Sin(rayAngle)*maxDepth), color.RGBA{0, 255, 0, 255})
	}
}

func collision(xPos, yPos int) bool {
	// Get 'tile' player is on
	x := ((xPos - 55) / (scrHeight / gMapGridHeight))
	y := ((yPos) / (scrHeight / gMapGridHeight))
	fmt.Printf("On tile [%d][%d]", x, y)
	fmt.Println()
	return gMap[y][x] == 1
}
