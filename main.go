package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player *Player
	count  int
}

const initialPositionX, initialPositionY = 340, 400

func (g *Game) Update() error {
	if g.player == nil {
		g.player = &Player{positionX: initialPositionX, positionY: initialPositionY, isOnFloor: true, animation: &Animation{lastSprite: 0}}
	}
	isIdle(g, 19)
	isHoldingJump(g, 5)
	isJumping(g, 3)

	g.count++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	position := &ebiten.DrawImageOptions{}
	position.GeoM.Translate(float64(g.player.positionX), float64(g.player.positionY))
	screen.DrawImage(g.player.animation.image, position)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Jeemp")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
