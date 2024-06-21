package game

import (
	"math/rand"

	"github.com/AtleSkaanes/gosnake/game/types"
)

var mainSnake types.Snake = types.NewSnake(types.NewVec2(0, 0), 0, types.Up)
var apple types.Vec2 = types.NewVec2(1, 1)

var score int
var gameConf types.GameConf

var IsInitialized bool = false

func Init(width uint8, height uint8, canLoop bool) {
	snakePos := types.NewVec2(3, 3)
	mainSnake = types.NewSnake(snakePos, 3, types.Right)

	score = 0
	dimensions := types.NewVec2(int(width), int(height))
	gameConf = types.GameConf{Dimensions: dimensions, CanLoop: canLoop}

	generateApple()

	IsInitialized = true
}

func Update(moveDir types.Direction) types.GameEvent {
	if moveDir != types.NoDirection && mainSnake.Direction != moveDir.GetOpposite() {
		mainSnake.Direction = moveDir
	}

	if !mainSnake.Move(gameConf) {
		return types.GameLost
	}

	if mainSnake.GetHead().IsEqual(apple) {
		score += 1
		mainSnake.Extend()
		generateApple()
		return types.AteApple
	}

	return types.GameContinue
}

func GetApple() types.Vec2 {
	return apple
}

func GetSnake() types.Snake {
	return mainSnake
}

func GetScore() int {
	return score
}

func GetConf() types.GameConf {
	return gameConf
}

func generateApple() {
	for {
		posX := rand.Intn(int(gameConf.Dimensions.X))
		posY := rand.Intn(int(gameConf.Dimensions.Y))
		pos := types.NewVec2(posX, posY)

		if !mainSnake.IsOnBody(pos) && !mainSnake.GetHead().IsEqual(pos) {
			apple = pos
			return
		}
	}
}
