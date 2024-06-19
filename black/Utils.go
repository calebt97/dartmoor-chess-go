package black

import (
	"github.com/notnil/chess"
	"sort"
)

type ScoredMove struct {
	move  *chess.Move
	score float32
}
type ByScore []ScoredMove

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].score > a[j].score }

func max(a, b float32) float32 {
	if a < b {
		return b
	}
	return a
}

func getBestMoves(game *chess.Game, listLength int) []*chess.Move {
	potentialMoves := game.ValidMoves()
	scoredMoves := make([]ScoredMove, len(potentialMoves))

	for i, move := range potentialMoves {

		next := game.Clone()

		if err := next.Move(move); err != nil {
			panic(err)
		}

		eval := evalBoard(next.Position())
		scoredMoves[i] = ScoredMove{
			move:  move,
			score: eval,
		}
	}
	sort.Sort(ByScore(scoredMoves))

	sortedMoves := make([]*chess.Move, listLength)

	size := int(min(float32(listLength), float32(len(potentialMoves))))

	for i := 0; i < size; i++ {
		sortedMoves[i] = scoredMoves[i].move
	}

	return sortedMoves
}

func min(a, b float32) float32 {
	if a > b {
		return b
	}
	return a
}
