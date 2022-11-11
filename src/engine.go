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
	tileX := 55 + (player.currentTileX * tileSize)
	tileY := player.currentTileY * tileSize

	for i := 0; i < numRays; i++ {
		sinA := math.Sin(rayAngle)
		cosA := math.Cos(rayAngle)

		// horizontals
		yHor := float64(tileY) - 1e-6

		dY := float64(-tileSize)
		if sinA > 0 {
			yHor = float64(tileY + tileSize)
			dY = float64(tileSize)
		}
		depthHor := (yHor - player.yPos) / sinA
		xHor := player.xPos + depthHor*cosA
		deltaDepth := float64(dY) / sinA
		dX := deltaDepth * cosA

		for d := 0; d < 20; d++ {
			// find which tile has been collided
			tileColX := int(((xHor - 55) / (scrHeight / gMapGridHeight)))
			tileColY := int(((yHor) / (scrHeight / gMapGridHeight)))

			fmt.Println(tileColY)
			if tileColX >= 0 && tileColY >= 0 && tileColY < 10 && tileColX < 15 && gMap[tileColY][tileColX] == 1 {
				break
			}
			xHor += float64(dX)
			yHor += float64(dY)
			depthHor += deltaDepth
		}

		// verticals

		// cos(angle) will be >0 if the ray intersects a vertical to
		// the right, otherwise it intersects to the left vertical
		xVert := float64(tileX) - float64(0.0000000001)
		dX = float64(-tileSize)
		if cosA > 0 {
			xVert = float64(tileX + tileSize)
			dX = float64(tileSize)
		}

		depthVert := (float64(xVert) - player.xPos) / cosA

		yVert := player.yPos + depthVert*sinA

		deltaDepth = float64(dX) / cosA
		dY = deltaDepth * sinA

		// do collision detection up to maximum depth
		for d := 0; d < 20; d++ {
			// find which tile has been collided
			if xVert >= 0 && yVert >= 0 {
				tileColX := int(((xVert - 55) / (scrHeight / gMapGridHeight)))
				tileColY := int(((yVert) / (scrHeight / gMapGridHeight)))

				if tileColY < 10 && tileColX < 15 && gMap[tileColY][tileColX] == 1 {
					break
				}
				xVert += float64(dX)
				yVert += float64(dY)
				depthVert += deltaDepth
			}

		}

		// depth
		depth := depthHor
		if depthVert < depthHor {
			depth = depthVert
		}

		rayAngle += deltaAngle

		ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(rayAngle)*float64(depth)), player.yPos+(math.Sin(rayAngle)*float64(depth)), color.RGBA{0, 255, 0, 255})
	}
}

func collision(xPos, yPos int) bool {
	// Get 'tile' player is on
	player.currentTileX = ((xPos - 55) / (scrHeight / gMapGridHeight))
	player.currentTileY = ((yPos) / (scrHeight / gMapGridHeight))
	return gMap[player.currentTileY][player.currentTileX] == 1
}
