package t3inverserules

import (
	"errors"
	"t3/t3board"
	"t3/t3utils"
	"t3/t3rules"
)

const (
	x = "X"
	o = "O"
	n = " "
)

type InverseTicTacToeGameRules struct {
	x             string
	o             string
	n             string
	defaultAction string
}

func NewInverseTicTacToeGameRules() *InverseTicTacToeGameRules {
	return &InverseTicTacToeGameRules{
		x: x,
		o: o,
		n: n,
	}
}

func (r *InverseTicTacToeGameRules) GetDefaultAction() string {
	return r.defaultAction
}

func (r *InverseTicTacToeGameRules) IsValidPlayer(player string) bool {
	return (player == x || player == o)
}

func (r *InverseTicTacToeGameRules) IsValidMove(board t3board.TicTacToeBoard, move *[2]int) (bool, error) {
	state, err := board.GetState(move)
	if err != nil {
		return false, err
	}
	return state == n, nil
}

func (r *InverseTicTacToeGameRules) IsValidTurn(board t3board.TicTacToeBoard, play string) bool {
	return play != board.GetLastPlay()
}

func (r *InverseTicTacToeGameRules) HasGameStarted(board t3board.TicTacToeBoard) bool {
	return t3utils.HasGameStarted(board.GetBoard())
}

func (r *InverseTicTacToeGameRules) HasGameEnded(board t3board.TicTacToeBoard) bool {
	return t3utils.HasGameEnded(board.GetBoard())
}

func (r *InverseTicTacToeGameRules) IsGameDraw(board t3board.TicTacToeBoard) bool {
	return t3utils.CheckIfDraw(board.GetBoard())
}

func (r *InverseTicTacToeGameRules) CanMakeMove(board t3board.TicTacToeBoard, play string, move [2]int, action string) (bool, error) {
	if !r.IsValidPlayer(play) {
		return false, errors.New("invalid player state")
	}
	isValidMove, err := r.IsValidMove(board, &move)
	if err != nil {
		return false, err
	}
	if !isValidMove {
		return false, errors.New("invalid move")
	}
	isValidTurn := r.IsValidTurn(board, play)
	if !isValidTurn {
		return false, errors.New("invalid turn")
	}
	if r.HasGameEnded(board) {
		return false, errors.New("game has ended")
	}
	return true, nil
}

func (r *InverseTicTacToeGameRules) GetResult(board t3board.TicTacToeBoard) t3rules.GameResult {
	result := t3rules.GameResult{
		IsEnded: false,
		IsDraw:  false,
		Winner:  "",
	}

	if !r.HasGameStarted(board) {
		return result
	}

	if r.IsGameDraw(board) {
		result.IsDraw = true
		result.IsEnded = true
		return result
	}

	// if game has ended and not draw then there is a winner
	if r.HasGameEnded(board) {
		result.IsEnded = true
		result.Winner = r.TogglePlay(board.GetLastPlay())
	}

	return result
}

func (r *InverseTicTacToeGameRules) Toss() string {
	return t3utils.Toss()
}

func (r *InverseTicTacToeGameRules) TogglePlay(play string) string {
	return t3utils.TogglePlay(play)
}

func (r *InverseTicTacToeGameRules) GetWinRow(board t3board.TicTacToeBoard) [3][2]int {
	return t3utils.GetWinRow(board.GetBoard(), board.GetLastPlay())
}