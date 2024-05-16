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

// def minimax(self, depth, board: chess.Board, maximizing_player,
// alpha, beta):
//
// MAX, MIN = 1000, -1000
// white_check_mate, black_checkmate = 1000, -1000
//
// if board.is_checkmate() and board.turn == chess.WHITE:
// # Bias for checkmates that are fewer turns away
// return black_checkmate + (depth * 2)
//
// if board.is_checkmate() and board.turn == chess.BLACK:
// # Bias for checkmates that are fewer turns away
// return white_check_mate - (depth * 2)
//
// if board.is_fivefold_repetition():
// return 0
//
// legal_moves = board.legal_moves
//
// # Terminating condition. i.e
// # leaf node is reached
// if depth >= self.standard_depth:
//
// if self.previous_evals.get(board.fen()) is not None:
// return self.previous_evals[board.fen()]
//
// eval = self.board_model.eval_board(board)
//
// self.previous_evals[board.fen()] = eval
//
// return eval
//
// if maximizing_player:
//
// best = MIN
//
// for move in legal_moves:
// val = 0
// board_copy = board.copy()
//
// q_search = self.needs_q_search(move, board_copy)
//
// board_copy.push(move)
//
// if q_search:
// val = self.qSearch(depth + 1, board_copy, False, alpha, beta)
//
// else:
// val = self.minimax(depth + 1, board_copy,
// False, alpha, beta)
//
// best = max(best, val)
// alpha = max(alpha, best)
//
// # Alpha Beta Pruning
// if beta <= alpha:
// break
//
// return best
//
// else:
// best = MAX
//
// for move in legal_moves:
// val = 0
// board_copy = board.copy()
//
// q_search = self.needs_q_search(move, board_copy)
//
// board_copy.push(move)
//
// if q_search:
// val = self.qSearch(depth + 1, board_copy, True, alpha, beta)
//
// else:
// val = self.minimax(depth + 1, board_copy,
// True, alpha, beta)
// best = min(best, val)
// beta = min(beta, best)
//
// # Alpha Beta Pruning
// if beta <= alpha:
// break
//
// return best
var SeenPositions = make(map[string]float32)
var NumPositions = 0

const (
	BlackCheckmate = -1000
	WhiteCheckmate = 1000
	MinEval        = -10000
	MaxEval        = 10000
	MaxDepth       = 3
	StartingAlpha  = -10000
	StartingBeta   = 10000
)

func FindMove(game *chess.Game) *chess.Move {
	var bestMove *chess.Move
	var bestEval *float32

	turn := game.Position().Turn()
	playerColor := turn == chess.White

	validMoves := game.ValidMoves()

	for _, move := range validMoves {
		next := game.Clone()

		if err := next.Move(move); err != nil {
			panic(err)
		}

		potential := minMax(0, StartingAlpha, StartingBeta, next, playerColor)

		if bestEval == nil || (potential < *bestEval && turn == chess.Black) {
			bestEval = &potential
			bestMove = move
		}

		if bestEval == nil || (potential > *bestEval && turn == chess.White) {
			bestEval = &potential
			bestMove = move
		}
	}
	log.Println("Best Move " + bestMove.String())
	log.Print("Num positions saved" + NumPositions)
	log.Print(*bestEval)

	return bestMove

}

func minMax(depth int, alpha float32, beta float32, game *chess.Game, maximizingPlayer bool) float32 {
	position := game.Position()

	if position.Turn() == chess.Black && position.Status() == chess.Checkmate {
		return BlackCheckmate + float32(depth*2)
	}

	if position.Turn() == chess.White && position.Status() == chess.Checkmate {
		return WhiteCheckmate - float32(depth*2)
	}
	if depth == MaxDepth {

		if value, exists := SeenPositions[game.FEN()]; exists {
			NumPositions++
			return value
		} else {
			value := evalBoard(game.Position())
			SeenPositions[game.FEN()] = value
			return value
		}
	}

	if maximizingPlayer {
		best := float32(MinEval)
		for _, move := range position.ValidMoves() {
			next := game.Clone()

			if err := next.Move(move); err != nil {
				panic(err)
			}

			value := minMax(depth+1, alpha, beta, next, false)
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

			if err := next.Move(move); err != nil {
				panic(err)
			}

			value := minMax(depth+1, alpha, beta, next, true)
			best = min(value, best)
			beta = min(beta, best)

			if beta >= alpha {
				break
			}
		}
		return best
	}

}

func max(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func min(a, b float32) float32 {
	if a > b {
		return b
	}
	return a
}
