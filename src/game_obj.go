package main

import (
	"bytes"
	"image"
	"io/ioutil"

	"github.com/hajimehoshi/ebiten/v2"
)

var wallGray *ebiten.Image
var wallBlue *ebiten.Image
var wallSkeleton *ebiten.Image

func loadImage() {
	wallGrayFile, _ := ioutil.ReadFile("../textures/wall_gray.png")
	wallBlueFile, _ := ioutil.ReadFile("../textures/wall_blue.png")
	wallSkelFile, _ := ioutil.ReadFile("../textures/wall_blue_skel.png")

	wallGrayImage, _, _ := image.Decode(bytes.NewReader(wallGrayFile))
	wallBlueImage, _, _ := image.Decode(bytes.NewReader(wallBlueFile))
	wallSkelImage, _, _ := image.Decode(bytes.NewReader(wallSkelFile))

	wallGray = ebiten.NewImageFromImage(wallGrayImage)
	wallBlue = ebiten.NewImageFromImage(wallBlueImage)
	wallSkeleton = ebiten.NewImageFromImage(wallSkelImage)
}
