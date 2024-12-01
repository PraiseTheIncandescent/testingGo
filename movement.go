package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func isIdle(game *Game, frameCount int) {
	if inpututil.KeyPressDuration(ebiten.KeySpace) == 0 {
		playerImage := mustLoadImage("./animations/sweepy-idle.png")
		PlayerSprite = convertToSprite(game, playerImage, frameCount)
	}
}

func isHoldingJump(game *Game, frameCount int) {
	if inpututil.KeyPressDuration(ebiten.KeySpace) > 0 {
		playerImage := mustLoadImage("./animations/sweepy-hold-jump.png")
		PlayerSprite = convertToSprite(game, playerImage, frameCount)
	}
}

func isJumping(game *Game, frameCount int) {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		playerImage := mustLoadImage("./animations/sweepy-jump.png")
		PlayerSprite = convertToSprite(game, playerImage, frameCount)
	}
}
