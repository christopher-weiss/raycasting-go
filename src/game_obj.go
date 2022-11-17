package main

import (
	"bytes"
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
)

// images
var wallGray *ebiten.Image
var wallBlue *ebiten.Image
var wallSkeleton *ebiten.Image
var objects *ebiten.Image
var ghostSprites *ebiten.Image

// sprites
var wellSprite Sprite

// objects
var gameObjects []GameObject
var ghostObj GameObject

type GameObject struct {
	xPos, yPos float64
	sprite     Sprite
}

type Sprite struct {
	sprite        image.Image
	width, height int
}

func loadImages() {
	wallGrayFile, _ := ioutil.ReadFile("textures/wall_gray.png")
	wallBlueFile, _ := ioutil.ReadFile("textures/wall_blue.png")
	wallSkelFile, _ := ioutil.ReadFile("textures/wall_blue_skel.png")
	objectsFile, _ := ioutil.ReadFile("sprites/objects.png")
	ghostSpritesFile, _ := ioutil.ReadFile("textures/ghost_sprites.png")

	wallGrayImage, _, _ := image.Decode(bytes.NewReader(wallGrayFile))
	wallBlueImage, _, _ := image.Decode(bytes.NewReader(wallBlueFile))
	wallSkelImage, _, _ := image.Decode(bytes.NewReader(wallSkelFile))
	objectsImage, _, _ := image.Decode(bytes.NewReader(objectsFile))
	ghostSpritesImage, _, _ := image.Decode(bytes.NewReader(ghostSpritesFile))

	wallGray = ebiten.NewImageFromImage(wallGrayImage)
	wallBlue = ebiten.NewImageFromImage(wallBlueImage)
	wallSkeleton = ebiten.NewImageFromImage(wallSkelImage)
	objects = ebiten.NewImageFromImage(objectsImage)
	ghostSprites = ebiten.NewImageFromImage(ghostSpritesImage)

	// load sprites
	wellSprite = Sprite{objects.SubImage(calculateObjectRect(7, 3)), 128, 128}
	ghostSpriteSheet := Sprite{ghostSprites, 25, 25}

	gameObjects = []GameObject{
		{320, 250, wellSprite},
		{310, 250, wellSprite},
		{300, 250, wellSprite},
		{290, 250, wellSprite},
		{280, 250, wellSprite},
		{270, 250, wellSprite},
		{260, 250, wellSprite},
		{250, 250, wellSprite},
	}

	ghostObj = GameObject{251, 255, ghostSpriteSheet}
}

func calculateObjectRect(xIndex, yIndex int) image.Rectangle {
	spriteSize := 128

	min := image.Point{(spriteSize * xIndex) + xIndex, (spriteSize * yIndex) + yIndex}
	max := image.Point{min.X + spriteSize, min.Y + spriteSize}

	return image.Rectangle{min, max}
}
