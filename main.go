package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	isIdle(g, 19)
	isHoldingJump(g, 5)
	isJumping(g, 3)
	return nil
}

var PlayerSprite *ebiten.Image

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(PlayerSprite, nil)
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
