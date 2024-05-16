package black

import (
	"github.com/notnil/chess"
)

const (
	queenValue  = 9.0
	bishopValue = 3.3
	knightValue = 3.2
	rookValue   = 5.0
	pawnValue   = 1
)

var pieceValues = map[chess.PieceType]float32{
	chess.Queen:  queenValue,
	chess.Bishop: bishopValue,
	chess.Rook:   rookValue,
	chess.Knight: knightValue,
	chess.Pawn:   pawnValue,
}

func evalBoard(position *chess.Position) float32 {
	board := position.Board()

	score := float32(0)
	squareMap := board.SquareMap()
	for _, value := range squareMap {

		if value.Color() == chess.Black {
			score -= pieceValues[value.Type()]
		} else {
			score += pieceValues[value.Type()]
		}
	}

	return score
}
