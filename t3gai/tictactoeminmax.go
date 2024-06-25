package t3gai

import (
	"math"
	"t3/t3utils"
)

type minMaxInput struct {
	board *[][]string
	depth int
	isMaximizingPlayer bool
	player string
	lastMove *[2]int	
}

func minMax(input minMaxInput) int {
	// if the game has started and not ended
	// if input.lastMove != nil && t3utils.HasGameEnded(input.board, input.player, *input.lastMove) {
		if input.lastMove != nil && t3utils.HasGameEnded(input.board) {
		// if the last move leads to victory 
		if t3utils.CheckIfWinMove(input.board, input.player, *input.lastMove) {
			if input.isMaximizingPlayer {
				return 1
			} else {
				// leads to loss or just another move
				return -1
			}
		} else {
			// leads to draw
			return 0
		}
	}
	possibleMoves := t3utils.GetPossibleMoves(input.board)

	if (input.isMaximizingPlayer) {
		bestScore := math.Inf(-1)
		for _, move := range possibleMoves {
			r, c := move[0], move[1]
			boardCopy := t3utils.CopyBoard(input.board)
			(*boardCopy)[r][c] = input.player
			score := float64(minMax(minMaxInput{
				board: boardCopy,
				depth: input.depth + 1,
				isMaximizingPlayer: false,
				player: input.player,
				lastMove: &move,
			}))
			bestScore = math.Max(bestScore, score)
		}
		return int(bestScore)
	} else {
		bestScore := math.Inf(1)
		for _, move := range possibleMoves {
			r, c := move[0], move[1]
			boardCopy := t3utils.CopyBoard(input.board)
			(*boardCopy)[r][c] = t3utils.TogglePlay(input.player)
			score := float64(minMax(minMaxInput{
				board: boardCopy,
				depth: input.depth + 1,
				isMaximizingPlayer: true,
				player: input.player,
				lastMove: &move,
			}))
			bestScore = math.Min(bestScore, score)
		}
		return int(bestScore)
	}
}

func CalculateMove(board *[][]string, player string) *[2]int {
	possibleMoves := t3utils.GetPossibleMoves(board)

	if len(possibleMoves) == 9 {
		move := t3utils.FirstMove()
		return &move
	}

	bestScore := math.Inf(-1)
	var bestMove *[2]int

	for _, move := range possibleMoves {
		copyBoard := t3utils.CopyBoard(board)
		r, c := move[0], move[1]
		(*copyBoard)[r][c] = player
		score := minMax(minMaxInput{
			board: copyBoard,
			depth: 0,
			isMaximizingPlayer: true,
			player: player,
			lastMove: &move,
		})

		if score > int(bestScore) {
			bestScore = float64(score) 
			bestMove = &move
		}
	}

	return bestMove
}