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
	gMapGridWidth  = 20
	gMapGridHeight = 15

	// player settings
	playerSpeed   = 2
	rotationSpeed = 0.05 // radians per tick
	startingPosX  = 200
	startingPosY  = 200
	startingAngle = 0

	// engine settings
	fov           = math.Pi / 4
	halfFov       = fov / 2
	numRays       = 250
	deltaAngle    = fov / numRays
	maxDepth      = 500
	scalingFactor = scrWidth / numRays
)
