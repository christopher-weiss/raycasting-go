package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	xPos float64
	yPos float64
}

var player = Player{200, 200}

func movePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if !collision(int(player.xPos+3), int(player.yPos)) {
			player.xPos += 3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if !collision(int(player.xPos-3), int(player.yPos)) {
			player.xPos -= 3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if !collision(int(player.xPos), int(player.yPos-3)) {
			player.yPos -= 3
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if !collision(int(player.xPos), int(player.yPos+3)) {
			player.yPos += 3
		}
	}
}
