package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	xPos         float64
	yPos         float64
	angle        float64 // in radians
	speed        float64
	currentTileX int
	currentTileY int
}

var player = Player{startingPosX, startingPosY, startingAngle, 3, 3, 4}

func movePlayer() {
	sinA := math.Sin(player.angle)
	cosA := math.Cos(player.angle)
	dx := 0.0
	dy := 0.0
	speed := player.speed
	speedSin := speed * sinA
	speedCos := speed * cosA

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if !collisionDetected(player.xPos+speedCos, player.yPos+speedSin) {
			dx += speedCos
			dy += speedSin
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if !collisionDetected(player.xPos-speedCos, player.yPos-speedSin) {
			dx += -speedCos
			dy += -speedSin

		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if !collisionDetected(player.xPos+speedSin, player.yPos-speedCos) {
			dx += speedSin
			dy += -speedCos
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if !collisionDetected(player.xPos-speedSin, player.yPos+speedCos) {
			dx += -speedSin
			dy += speedCos
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.angle -= rotationSpeed
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.angle += rotationSpeed
	}

	// move player according input
	player.xPos += dx
	player.yPos += dy

	// set current tile indexes player is currently on
	player.currentTileX = ((int(player.xPos) - xOffset) / (scrHeight / gMapGridHeight))
	player.currentTileY = (int(player.yPos)) / (scrHeight / gMapGridHeight)
}
