package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func getMaxHeightJump(holdDuration int) int {
	if holdDuration <= 50 {
		return 300
	}
	if holdDuration <= 100 {
		return 250
	}
	return 150
}

func isIdle(game *Game, frameCount int) {
	if game.player.isOnFloor && inpututil.KeyPressDuration(ebiten.KeySpace) == 0 {
		playerImage := mustLoadImage("./animations/sweepy-idle.png")
		game.player.animation.isFinished = false
		game.player.animation.image = convertToSprite(game, playerImage, frameCount)
	}
}

func isHoldingJump(game *Game, frameCount int) {
	if game.player.isOnFloor && inpututil.KeyPressDuration(ebiten.KeySpace) > 0 {
		log.Output(1, ""+fmt.Sprint(inpututil.KeyPressDuration(ebiten.KeySpace)))
		game.player.holdDuration = inpututil.KeyPressDuration(ebiten.KeySpace)
		if !game.player.animation.isFinished {
			playerImage := mustLoadImage("./animations/sweepy-hold-jump.png")
			game.player.animation.image = convertToSprite(game, playerImage, frameCount)
		}
	}
}

func isJumping(game *Game, frameCount int) {
	if inpututil.IsKeyJustReleased(ebiten.KeySpace) {
		game.player.isOnFloor = false
	}

	if !game.player.isOnFloor {
		if !game.player.animation.isFinished {
			playerImage := mustLoadImage("./animations/sweepy-jump.png")
			game.player.animation.image = convertToSprite(game, playerImage, frameCount)
		}

		if !game.player.reachedTopOnJump {
			if game.player.positionY > getMaxHeightJump(game.player.holdDuration) {

				game.player.positionY -= 3
			} else {
				game.player.reachedTopOnJump = true
			}
		} else if game.player.reachedTopOnJump {
			if game.player.positionY < initialPositionY {
				game.player.positionY += 3
			} else {
				game.player.isOnFloor = true
				game.player.reachedTopOnJump = false
			}
		}
	}
}
