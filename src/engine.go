package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var f1PressedInLastTick = false
var showMap = false

func debugDraw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.0f", ebiten.ActualFPS()))

	// Draw Ceiling
	ebitenutil.DrawRect(screen, 0, 0, scrWidth, halfHeight, color.RGBA{30, 30, 30, 255})

	// Draw floor
	ebitenutil.DrawRect(screen, 0, halfHeight, scrWidth, scrHeight, color.RGBA{100, 100, 100, 255})

	// Draw game map
	tileSize := scrHeight / gMapGridHeight
	if showMap {
		xPadding := xOffset // x-axis padding to account for wide window
		for i := 0; i < gMapGridHeight; i++ {
			for j := 0; j < gMapGridWidth; j++ {
				if gMap[i][j] >= 1 {
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
		ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(player.angle)), player.yPos+(math.Sin(player.angle)), color.RGBA{0, 0, 255, 60})
	}

	// Draw rays
	rayAngle := player.angle - halfFov + 0.0001
	tileX := xOffset + (player.currentTileX * tileSize)
	tileY := player.currentTileY * tileSize

	for rayNum := 0; rayNum < numRays; rayNum++ {
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
			if collisionDetected(xHor, yHor) >= 1 {
				break
			}
			xHor += float64(dX)
			yHor += float64(dY)
			depthHor += deltaDepth
		}

		// verticals

		// cos(angle) will be > 0 if the ray intersects a vertical to
		// the right, otherwise it intersects to the left vertical
		xVert := float64(tileX) - 1e-6
		dX = float64(-tileSize)
		if cosA > 0 {
			xVert = float64(tileX + tileSize)
			dX = float64(tileSize)
		}

		depthVert := (float64(xVert) - player.xPos) / cosA

		yVert := player.yPos + depthVert*sinA
		if yVert < 0 {
			yVert = 0
		}

		deltaDepth = float64(dX) / cosA
		dY = deltaDepth * sinA

		// do collision detection up to maximum depth
		for d := 0; d < 50; d++ {
			if collisionDetected(xVert, yVert) >= 1 {
				break
			}
			xVert += float64(dX)
			yVert += float64(dY)

			depthVert += deltaDepth

		}

		// depth calculation of ray until closest collision
		depth := depthHor
		verticalCollision := false

		if depthVert < depthHor {
			depth = depthVert
			verticalCollision = true
		}

		// correct fishbowl distortion
		depthFaCorrected := depth * (math.Cos(player.angle - rayAngle))

		// projection
		screenDistance := float64(halfWidth) / math.Tan(halfFov)
		projHeight := screenDistance * wallHeight / (depthFaCorrected + 0.0001)

		// texture mapping
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, projHeight/(textureSize))
		op.GeoM.Translate(float64(rayNum*scalingFactor), halfHeight-(math.Floor(projHeight/2)))
		op.Filter = ebiten.FilterLinear

		brightness := 1 / (1 + math.Pow(depth, 5.0)*0.000000001)
		op.ColorM.ChangeHSV(0, 1, brightness)

		xCoord := 0.0
		yCoord := 0.0
		if verticalCollision {
			xCoord = xVert
			yCoord = yVert
		} else {
			xCoord = xHor
			yCoord = yHor
		}
		// calculate subImage offset
		var min, max image.Point

		// since the yVert/xHor Values can be identical after casting to int, this value stretches
		// out the values so that texture mapping works correctly.
		stretchFactor := 14.7
		if verticalCollision {
			if cosA > 0 {
				min = image.Point{(int(yVert*stretchFactor) % textureSize), 0}
			} else {
				min = image.Point{textureSize - int(yVert*stretchFactor)%textureSize, 0}
			}
		} else {
			if sinA > 0 {
				min = image.Point{textureSize - int(xHor*stretchFactor)%textureSize, 0}
			} else {
				min = image.Point{int(xHor*stretchFactor) % textureSize, 0}
			}
		}

		if min.X+scalingFactor >= textureSize {
			min.X = textureSize - scalingFactor
		}
		max = image.Point{min.X + scalingFactor, textureSize}

		rect := image.Rectangle{min, max}

		if showMap {
			ebitenutil.DrawCircle(screen, xCoord, yCoord, 1, color.RGBA{0, 255, 0, 60})
		}
		if collisionDetected(xCoord, yCoord) == 1 {
			screen.DrawImage(wallGray.SubImage(rect).(*ebiten.Image), op)
		}
		if collisionDetected(xCoord, yCoord) == 2 {
			screen.DrawImage(wallBlue.SubImage(rect).(*ebiten.Image), op)
		}
		if collisionDetected(xCoord, yCoord) == 3 {
			screen.DrawImage(wallSkeleton.SubImage(rect).(*ebiten.Image), op)
		}

		rayAngle += deltaAngle

		if showMap {
			if collisionDetected(xVert, yVert) != -1 {
				rayColor := color.RGBA{0, 255, 0, 20}
				ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(rayAngle)*float64(depth)), player.yPos+(math.Sin(rayAngle)*float64(depth)), rayColor)
			} else {
				ebitenutil.DrawLine(screen, player.xPos, player.yPos, player.xPos+(math.Cos(rayAngle)*float64(depth)), player.yPos+(math.Sin(rayAngle)*float64(depth)), color.RGBA{255, 0, 0, 20})
			}
		}
	}

	// We need to sort the objects to be displayed according to the viewing angle,
	// so that objects that are farther away are not drawn over nearer objects
	sort.Slice(gameObjects, func(p, q int) bool {
		if math.Cos(player.angle) > 0 {
			return gameObjects[p].xPos > gameObjects[q].xPos
		} else {
			return gameObjects[p].xPos < gameObjects[q].xPos
		}
	})

	for sp := 0; sp < len(gameObjects); sp++ {
		drawStaticSprite(gameObjects[sp], screen)
	}
	drawAnimatedSprite(ghostObj, screen)

	if showMap {
		ebitenutil.DrawCircle(screen, 250, 250, 1, color.White)
		ebitenutil.DrawCircle(screen, 260, 260, 1, color.RGBA{255, 0, 0, 255})
	}
}

func drawStaticSprite(obj GameObject, screen *ebiten.Image) {
	// calculate delta angle, which is the angle between the players line of sight
	// and the angle of the plane of the player and the object (theta).
	// p---->
	// |\`
	// | \ `
	// |  \ ∆` line-of-sight
	// v   obj
	theta := math.Atan2((obj.yPos - player.yPos), (obj.xPos - player.xPos))
	delta := theta - player.angle

	// o
	if (obj.xPos > 0 && player.angle > math.Pi) || (obj.xPos < 0 && obj.yPos < 0) {
		delta += math.Pi * 2
	}

	deltaRays := delta / deltaAngle

	scrXPos := (numRays/2 + deltaRays) * scalingFactor
	dist := math.Hypot(obj.xPos-player.xPos, obj.yPos-player.yPos)
	normalizedDist := dist * math.Cos(delta)

	screenDistance := float64(halfWidth) / math.Tan(halfFov)
	proj := (screenDistance / normalizedDist) / 100
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(proj, proj)
	options.GeoM.Translate(scrXPos-float64(obj.sprite.width/2), float64(halfHeight)-float64(obj.sprite.height/2)-((proj*100)/2)+70)
	screen.DrawImage(obj.sprite.sprite.(*ebiten.Image), options)
}

func drawAnimatedSprite(obj GameObject, screen *ebiten.Image) {
	theta := math.Atan2(obj.yPos-player.yPos, obj.xPos-player.xPos)
	delta := theta - player.angle
	if (obj.xPos > 0 && player.angle > math.Pi) || (obj.xPos < 0 && obj.yPos < 0) {
		delta += math.Pi * 2
	}
	deltaRays := delta / deltaAngle
	scrXPos := (numRays/2 + deltaRays) * scalingFactor
	dist := math.Hypot(obj.xPos-player.xPos, obj.yPos-player.yPos)
	normalizedDist := dist * math.Cos(delta)

	screenDistance := float64(halfWidth) / math.Tan(halfFov)
	proj := ((screenDistance / normalizedDist) / 100) + 0.2
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(proj, proj)
	options.GeoM.Translate(scrXPos-float64(obj.sprite.width/2), float64(halfHeight)-float64(obj.sprite.height/2)-((proj*10)/2)+20)

	spriteY := ((tick / 10) % 4) * obj.sprite.height

	screen.DrawImage(ghostSprites.SubImage(image.Rectangle{image.Point{0, spriteY}, image.Point{obj.sprite.width, spriteY + obj.sprite.height}}).(*ebiten.Image), options)
}

func collisionDetected(xPos, yPos float64) int {
	// Get indexes of the tile
	tileColX := int(((xPos - xOffset) / (scrHeight / gMapGridHeight)))
	tileColY := int(((yPos) / (scrHeight / gMapGridHeight)))

	if tileColX < 0 || tileColX >= gMapGridWidth || tileColY < 0 || tileColY >= gMapGridHeight {

		if tileColX <= 0 {
			tileColX = 1
		}
		if tileColX >= gMapGridWidth {
			tileColX = gMapGridWidth - 1
		}
		if tileColY <= 0 {
			tileColY = 1
		}
		if tileColY >= gMapGridHeight {
			tileColY = gMapGridHeight - 1
		}
	}

	return gMap[tileColY][tileColX]
}
