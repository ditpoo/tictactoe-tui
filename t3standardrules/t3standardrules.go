package t3standardrules

import (
	"errors"
	"t3/t3board"
	"t3/t3rules"
	"t3/t3utils"
)

const (
	x = "X"
	o = "O"
	n = " "
)

type StandardTicTacToeGameRules struct {
	x             string
	o             string
	n             string
	defaultAction string
}

func NewStandardTicTacToeGameRules() *StandardTicTacToeGameRules {
	return &StandardTicTacToeGameRules{
		x: x,
		o: o,
		n: n,
	}
}

func (r *StandardTicTacToeGameRules) GetDefaultAction() string {
	return r.defaultAction
}

func (r *StandardTicTacToeGameRules) IsValidPlayer(player string) bool {
	return (player == x || player == o)
}

func (r *StandardTicTacToeGameRules) IsValidMove(board t3board.TicTacToeBoard, move *[2]int) (bool, error) {
	state, err := board.GetState(move)
	if err != nil {
		return false, err
	}
	return state == n, nil
}

func (r *StandardTicTacToeGameRules) IsValidTurn(board t3board.TicTacToeBoard, play string) bool {
	return play != board.GetLastPlay()
}

func (r *StandardTicTacToeGameRules) HasGameStarted(board t3board.TicTacToeBoard) bool {
	return t3utils.HasGameStarted(board.GetBoard())
}

func (r *StandardTicTacToeGameRules) HasGameEnded(board t3board.TicTacToeBoard) bool {
	return t3utils.HasGameEnded(board.GetBoard())
}

func (r *StandardTicTacToeGameRules) IsGameDraw(board t3board.TicTacToeBoard) bool {
	return t3utils.CheckIfDraw(board.GetBoard())
}

func (r *StandardTicTacToeGameRules) CanMakeMove(board t3board.TicTacToeBoard, play string, move [2]int, action string) (bool, error) {
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

func (r *StandardTicTacToeGameRules) GetResult(board t3board.TicTacToeBoard) t3rules.GameResult {
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
		result.Winner = board.GetLastPlay()
	}

	return result
}

func (r *StandardTicTacToeGameRules) Toss() string {
	return t3utils.Toss()
}

func (r *StandardTicTacToeGameRules) TogglePlay(play string) string {
	return t3utils.TogglePlay(play)
}

func (r *StandardTicTacToeGameRules) GetWinRow(board t3board.TicTacToeBoard) [3][2]int {
	return t3utils.GetWinRow(board.GetBoard(), board.GetLastPlay())
}