package black

import (
	"github.com/notnil/chess"
)

const (
	queenValue    = 9.0
	bishopValue   = 3.3
	knightValue   = 3.2
	rookValue     = 5.0
	pawnValue     = 1
	castlingBonus = 3.0
)

var pieceValues = map[chess.PieceType]float32{
	chess.Queen:  queenValue,
	chess.Bishop: bishopValue,
	chess.Rook:   rookValue,
	chess.Knight: knightValue,
	chess.Pawn:   pawnValue,
}

var score = float32(0)

func reverseIndex(index int) int {
	return 63 - index
}

func evalBoard(position *chess.Position) float32 {
	score = float32(0)

	checkCastling(position)
	handlePieceMap(position)

	return score
}

func checkCastling(position *chess.Position) {

	if position.CastleRights().CanCastle(chess.Black, chess.KingSide) {
		score -= castlingBonus
	}

	if position.CastleRights().CanCastle(chess.White, chess.KingSide) {
		score += castlingBonus
	}

}

func handlePieceMap(position *chess.Position) {

	squareMap := position.Board().SquareMap()

	for square, piece := range squareMap {

		toBeAdded := float32(0)
		pieceType := piece.Type()
		index := int(square)

		if piece.Color() == chess.Black {
			index = reverseIndex(index)
		}

		switch pieceType {
		case chess.Pawn:
			toBeAdded += pawnTable[index]
		case chess.Knight:
			toBeAdded += knightTable[index] / 20
		case chess.Bishop:
			toBeAdded += bishopTable[index] / 20
		case chess.Rook:
			toBeAdded += rookTable[index] / 20
		case chess.Queen:
			toBeAdded += queenTable[index] / 20
		case chess.King:
			toBeAdded += kingTable[index] / 20
		}

		toBeAdded += pieceValues[pieceType] * 10

		if piece.Color() == chess.Black {
			toBeAdded *= -1
		}

		score += toBeAdded
	}

}
