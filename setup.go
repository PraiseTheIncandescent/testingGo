package main

import (
	"bytes"
	"image"
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 720
	screenHeight = 480

	frameOX     = 0
	frameOY     = 0
	frameWidth  = 64
	frameHeight = 64
)

func mustLoadImage(name string) *ebiten.Image {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Lee el archivo en un slice de bytes
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader([]byte(fileBytes)))

	if err != nil {
		log.Fatal(err)
	}

	return ebiten.NewImageFromImage(img)
}

func convertToSprite(g *Game, playerImage *ebiten.Image, frameCount int) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 8) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY

	return playerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
}
