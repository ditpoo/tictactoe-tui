package t3neutralrules

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

type NeutralTicTacToeGameRules struct {
	x             string
	o             string
	n             string
	defaultAction string
}

func NewNeutralTicTacToeGameRules() *NeutralTicTacToeGameRules {
	return &NeutralTicTacToeGameRules{
		x: x,
		o: o,
		n: n,
	}
}

func (r *NeutralTicTacToeGameRules) GetDefaultAction() string {
	return r.defaultAction
}

func (r *NeutralTicTacToeGameRules) IsValidPlayer(player string) bool {
	return (player == x || player == o)
}

func (r *NeutralTicTacToeGameRules) IsValidMove(board t3board.TicTacToeBoard, move *[2]int) (bool, error) {
	state, err := board.GetState(move)
	if err != nil {
		return false, err
	}
	return state == n, nil
}

func (r *NeutralTicTacToeGameRules) IsValidTurn(board t3board.TicTacToeBoard, play string) bool {
	return play != board.GetLastPlay()
}

func (r *NeutralTicTacToeGameRules) HasGameStarted(board t3board.TicTacToeBoard) bool {
	return t3utils.HasGameStarted(board.GetBoard())
}

func (r *NeutralTicTacToeGameRules) HasGameEnded(board t3board.TicTacToeBoard) bool {
	return t3utils.HasGameEnded(board.GetBoard())
}

func (r *NeutralTicTacToeGameRules) IsGameDraw(board t3board.TicTacToeBoard) bool {
	return t3utils.CheckIfDraw(board.GetBoard())
}

func (r *NeutralTicTacToeGameRules) CanMakeMove(board t3board.TicTacToeBoard, play string, move [2]int, action string) (bool, error) {
	isValidMove, err := r.IsValidMove(board, &move)
	if err != nil {
		return false, err
	}
	if !isValidMove {
		return false, errors.New("invalid move")
	}
	if r.HasGameEnded(board) {
		return false, errors.New("game has ended")
	}
	return true, nil
}

func (r *NeutralTicTacToeGameRules) MakeMove(board *t3board.TicTacToeBoard, play string, move [2]int) (bool, error) {
	isStateSet, err := board.SetState(&move, x)
	if err != nil {
		return false, err
	}
	if isStateSet {
		board.SetLastPlay(play)
		board.SetlastMove(&move)
	}
	return isStateSet, err
}

func (r *NeutralTicTacToeGameRules) GetResult(board t3board.TicTacToeBoard) t3rules.GameResult {
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

func (r *NeutralTicTacToeGameRules) Toss() string {
	return t3utils.Toss()
}

func (r *NeutralTicTacToeGameRules) TogglePlay(play string) string {
	return t3utils.TogglePlay(play)
}

func (r *NeutralTicTacToeGameRules) GetWinRow(board t3board.TicTacToeBoard) [3][2]int {
	return t3utils.GetWinRow(board.GetBoard(), board.GetLastPlay())
}