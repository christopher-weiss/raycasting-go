package main

import "math"

const (
	// window & screen settings
	winWidth   = 1600
	winHeight  = 900
	scrWidth   = winWidth / 2
	scrHeight  = winHeight / 2
	halfWidth  = scrWidth / 2
	halfHeight = scrHeight / 2

	xOffset = 0

	// map settings
	gMapGridWidth  = 20
	gMapGridHeight = 15

	// player settings
	playerSpeed   = 0.4
	rotationSpeed = 0.025 // radians per tick
	startingPosX  = 200
	startingPosY  = 200
	startingAngle = 0

	// engine settings
	fov           = math.Pi / 3
	halfFov       = fov / 2
	numRays       = 400
	deltaAngle    = fov / numRays
	maxDepth      = 1000
	scalingFactor = scrWidth / numRays
	wallHeight    = 4

	// texture settings
	textureSize     = 65
	halfTextureSize = textureSize / 2
)
