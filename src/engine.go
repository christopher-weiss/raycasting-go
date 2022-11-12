package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var showMapPressed = false
var showMap = false

func debugDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.0f", ebiten.ActualFPS()))

	// Draw game map
	tileSize := scrHeight / gMapGridHeight
	if showMap {
		xPadding := 55 // x-axis padding to account for wide window
		for i := 0; i < gMapGridHeight; i++ {
			for j := 0; j < gMapGridWidth; j++ {
				if gMap[i][j] == 1 {
					ebitenutil.DrawRect(screen, float64(xPadding+(j*tileSize)), float64(i*tileSize), float64(tileSize-1), float64(tileSize-1), color.RGBA{255, 255, 255, 60})
				}
			}
		}
	}

	// Draw player
	if showMap {
		ebitenutil.DrawCircle(screen, player.xPos, player.yPos, 3, color.RGBA{255, 0, 0, 60})
	}

	// Draw line of sight
	if showMap {
		ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(player.angle)*20), player.yPos+(math.Sin(player.angle)*20), color.RGBA{0, 0, 255, 60})
	}

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
			if collisionDetected(xHor, yHor) {
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
				if collisionDetected(xVert, yVert) {
					break
				}
				xVert += float64(dX)
				yVert += float64(dY)
				depthVert += deltaDepth
			}

		}

		// depth calculation of ray until closest collision
		depth := depthHor
		if depthVert < depthHor {
			depth = depthVert
		}

		// projection
		screenDistance := float64(halfWidth) / math.Tan(halfFov)
		projHeight := screenDistance / (depth + 0.0001)

		depthColor := 255 / 1
		ebitenutil.DrawRect(screen, float64(i*scalingFactor), halfHeight-(projHeight/2), scalingFactor, projHeight, color.RGBA{255, 255, 255, uint8(depthColor)})

		rayAngle += deltaAngle

		if showMap {
			ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(rayAngle)*float64(depth)), player.yPos+(math.Sin(rayAngle)*float64(depth)), color.RGBA{0, 255, 0, 20})
		}
	}
}

func collisionDetected(xPos, yPos float64) bool {
	// Get indexes of the tile
	tileColX := int(((xPos - xOffset) / (scrHeight / gMapGridHeight)))
	tileColY := int(((yPos) / (scrHeight / gMapGridHeight)))

	if tileColX < 0 || tileColX >= gMapGridWidth || tileColY < 0 || tileColY >= gMapGridHeight {
		return false
	}

	return gMap[tileColY][tileColX] == 1
}
