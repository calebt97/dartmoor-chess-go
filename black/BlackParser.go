package black

import (
	"github.com/notnil/chess"
	"log"
)

//Randotron
//func FindMove(game *chess.Game) *chess.Move {
//	validMoves := game.ValidMoves()
//	log.Print(game.Position().Board().SquareMap())
//
//	if len(valid) == 0 {
//		return nil
//	}
//
//	return valid[rand.Intn(len(valid))]
//}

const (
	BlackCheckmate = -1000
	WhiteCheckmate = 1000
	MinEval        = -10000
	MaxEval        = 10000
	MaxDepth       = 4
	StartingAlpha  = -10000
	StartingBeta   = 10000
)

var SeenPositions = make(map[[16]byte]float32)

func FindMove(game *chess.Game) *chess.Move {
	var bestMove *chess.Move
	var bestEval *float32

	turn := game.Position().Turn()
	playerColor := turn == chess.White

	validMoves := getBestMoves(game, 20)

	for _, move := range validMoves {
		next := game.Clone()

		if err := next.Move(move); err != nil {
			continue
		}

		potential := minMax(0, StartingAlpha, StartingBeta, next, playerColor)

		if bestEval == nil || (potential < *bestEval && turn == chess.Black) {
			//bestGame = next
			bestEval = &potential
			bestMove = move
		}

		if bestEval == nil || (potential > *bestEval && turn == chess.White) {
			//bestGame = next
			bestEval = &potential
			bestMove = move
		}
	}
	log.Println("Best Move " + bestMove.String())
	log.Println(*bestEval)

	return bestMove

}

func minMax(depth int, alpha float32, beta float32, game *chess.Game, maximizingPlayer bool, qValue int) float32 {
	position := game.Position()

	if position.Turn() == chess.White && position.Status() == chess.Checkmate {
		return BlackCheckmate + float32(depth*2)
	}

	if position.Turn() == chess.Black && position.Status() == chess.Checkmate {
		return WhiteCheckmate - float32(depth*2)
	}

	if depth >= MaxDepth || qValue <= 3 {

		if value, exists := SeenPositions[position.Hash()]; exists {
			return value
		} else {

			value := evalBoard(position)
			SeenPositions[position.Hash()] = value
			return value
		}

	}

	potentials := game.ValidMoves()

	if maximizingPlayer {
		best := float32(MinEval)
		for _, move := range potentials {
			next := game.Clone()
			value := float32(0)
			if position.Board().Piece(move.S2()).String() != nil {
				value = minMax(depth, alpha, beta, next, false, qValue+1)
			}

			if err := next.Move(move); err != nil {
				panic(err)
			}

			value = minMax(depth+1, alpha, beta, next, false, 0)
			best = max(value, best)
			alpha = max(alpha, best)

			if beta <= alpha {
				break
			}

		}
		return best
	} else {
		best := float32(MaxEval)
		for _, move := range position.ValidMoves() {
			next := game.Clone()
			value := float32(0)

			if err := next.Move(move); err != nil {
				panic(err)
			}

			value = minMax(depth+1, alpha, beta, next, true, 0)
			best = min(value, best)
			beta = min(beta, best)

			if beta >= alpha {
				break
			}
		}
		return best
	}

}
