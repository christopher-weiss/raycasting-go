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

	xOffset = 55

	// map settings
	gMapGridWidth  = 15
	gMapGridHeight = 10

	// player settings
	playerSpeed   = 0.4
	rotationSpeed = 0.03 // radians per tick
	startingPosX  = 200
	startingPosY  = 200
	startingAngle = 0

	// engine settings
	fov           = math.Pi / 3
	halfFov       = fov / 2
	numRays       = 200
	deltaAngle    = fov / numRays
	maxDepth      = 500
	scalingFactor = scrWidth / numRays
)
