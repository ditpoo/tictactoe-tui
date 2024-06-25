package t3game

import (
	"errors"
	"t3/t3board"
	"t3/t3rules"
)

const (
	x = "X"
	o = "O"
	n = " "
)

type move struct {
	play     string
	position [2]int
}

type TicTacToeGame struct {
	x           string
	o           string
	n           string
	tboard      t3board.TicTacToeBoard
	grules      t3rules.GameRulesManager
	moveHistory []move
}

func NewTicTacToeGame(board t3board.TicTacToeBoard, rules t3rules.GameRulesManager) (*TicTacToeGame, error) {
	return &TicTacToeGame{
		x:      x,
		o:      o,
		n:      n,
		tboard: board,
		grules: rules,
	}, nil
}

func (t *TicTacToeGame) GetLastPlay() string {
	return t.tboard.GetLastPlay()
}

func (t *TicTacToeGame) GetLastMove() [2]int {
	return t.tboard.GetlastMove()
}

func (t *TicTacToeGame) validatePlayer(player string) bool {
	return (player == x || player == o)
}

// this method assumes that it is called afer move has played and persisted in lastPlay and lastMove
func (t *TicTacToeGame) updateHistory() {
	t.moveHistory = append(t.moveHistory, move{play: t.tboard.GetLastPlay(), position: t.tboard.GetlastMove()})
}

func (t *TicTacToeGame) GetState(position *[2]int) (string, error) {
	return t.tboard.GetState(position)
}

func (t *TicTacToeGame) GetBoard() *[][]string {
	return t.tboard.GetBoard()
}

func (t *TicTacToeGame) SetMove(player string, position *[2]int) (bool, error) {
	if !t.validatePlayer(player) {
		return false, errors.New("invalid player state")
	}
	canMakeMove, err := t.grules.CanMakeMove(t.tboard, player, *position, "set") 
	if !canMakeMove {
		if err != nil {
			return false, err
		} else {
			return false, errors.New("can't make move")
		}
	}
	if player == t.tboard.GetLastPlay() {
		return false, errors.New("player can't make move two times in a row")
	}
	state, err := t.tboard.GetState(position)
	if err != nil {
		return false, err
	}
	if state != t.n {
		return false, errors.New("can't make move for that position")
	}
	isStateSet, err := t.tboard.SetState(position, player)
	if !isStateSet {
		if err != nil {
			return false, err
		} else {
			return false, errors.New("failed to set state or make move for that position")
		}
	}
	t.updateHistory()
	return true, nil
}

func (t *TicTacToeGame) HasGameStarted() bool {
	return t.grules.HasGameStarted(t.tboard)
}

func (t *TicTacToeGame) HasGameEnded() bool {
	return t.grules.HasGameEnded(t.tboard)
}

func (t *TicTacToeGame) GetResult() t3rules.GameResult {
	return t.grules.GetResult(t.tboard)
}

func (t *TicTacToeGame) GetRules() t3rules.GameRulesManager {
	return t.grules
}

func (t *TicTacToeGame) GetGameHistory() []move {
	return t.moveHistory
}

func (t *TicTacToeGame) Toss() string {
	return t.grules.Toss()
}

func (t *TicTacToeGame) TogglePlay() string {
	return t.grules.TogglePlay(t.GetLastPlay())
}