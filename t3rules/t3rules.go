package t3rules

import (
	"t3/t3board"
)

type GameResult struct {
	IsEnded bool
	IsDraw  bool
	Winner  string
}

type GameRulesManager interface {
	CanMakeMove(board t3board.TicTacToeBoard, play string, move [2]int, action string) (bool, error)
	GetDefaultAction() string
	GetResult(board t3board.TicTacToeBoard) GameResult
	GetWinRow(board t3board.TicTacToeBoard) [3][2]int
	HasGameStarted(board t3board.TicTacToeBoard) bool
	HasGameEnded(board t3board.TicTacToeBoard) bool
	IsValidMove(board t3board.TicTacToeBoard, move *[2]int) (bool, error)
	IsValidTurn(board t3board.TicTacToeBoard, play string) bool
	MakeMove(board *t3board.TicTacToeBoard, play string, move [2]int) (bool, error)
	Toss() string
	TogglePlay(play string) string
}