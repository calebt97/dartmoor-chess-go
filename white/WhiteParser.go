package white

import (
	"github.com/notnil/chess"
	"math/rand"
)

func FindMove(game *chess.Game) *chess.Move {
	valid := game.ValidMoves()
	if len(valid) == 0 {
		return nil
	}

	return valid[rand.Intn(len(valid))]
}
