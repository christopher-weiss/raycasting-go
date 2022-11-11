package main

import "math"

const (
	// window & screen settings
	winWidth  = 1600
	winHeight = 900
	scrWidth  = winWidth / 2
	scrHeight = winHeight / 2
	xOffset   = 55

	// map settings
	gMapGridWidth  = 15
	gMapGridHeight = 10

	// player settings
	movementSpeed = 3   // pixels per tick
	rotationSpeed = 0.1 // radians per tick
	startingPosX  = 200
	startingPosY  = 200
	startingAngle = 0

	// engine settings
	fov        = math.Pi / 3
	halfFov    = fov / 2
	numRays    = 100
	deltaAngle = fov / numRays
	maxDepth   = 500
)
