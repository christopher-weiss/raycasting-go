package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	xPos  float64
	yPos  float64
	angle float64 // in radians
	speed float64
	dx    float64
	dy    float64
}

var player = Player{200, 200, 0, 3, 0, 0}

func movePlayer() {
	sinA := math.Sin(player.angle)
	cosA := math.Cos(player.angle)
	dx := 0.0
	dy := 0.0
	speed := player.speed
	speedSin := speed * sinA
	speedCos := speed * cosA

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		if !collision(int(player.xPos+speedCos), int(player.yPos+speedSin)) {
			dx += speedCos
			dy += speedSin
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		if !collision(int(player.xPos-speedCos), int(player.yPos-speedSin)) {
			dx += -speedCos
			dy += -speedSin

		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		if !collision(int(player.xPos+speedSin), int(player.yPos-speedCos)) {
			dx += speedSin
			dy += -speedCos
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		if !collision(int(player.xPos-speedSin), int(player.yPos+speedCos)) {
			dx += -speedSin
			dy += speedCos
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		player.angle -= 0.1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		player.angle += 0.1
	}

	player.xPos += dx
	player.yPos += dy
	player.dx = dx
	player.dy = dy
}
