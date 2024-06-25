package t3board

import (
	"errors"
	"t3/board"
)

const (
	x = "X"
	o = "O"
	n = " "
)

type TicTacToeBoard struct {
	x        string
	o        string
	n        string
	board    board.BoardHandler
	lastPlay string
	lastMove [2]int
}

func setBoard(b board.BoardHandler, state string) error {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			isStateState, err := b.SetState(&[2]int{int(i), int(j)}, state)
			if !isStateState {
				if err != nil {
					return errors.New("failed to set zero state for board")
				}
			}
		}
	}
	return nil
}

func NewTicTacToeBoard(board board.BoardHandler) (*TicTacToeBoard, error) {
	size := board.GetSize()
	if size[0] != 3 && size[1] != 3 {
		return nil, errors.New("invalid row and column size, it should be 3")
	}
	err := setBoard(board, n)
	if err != nil {
		return nil, err
	}
	return &TicTacToeBoard{
		x:     x,
		o:     o,
		n:     n,
		board: board,
	}, nil
}

func (t *TicTacToeBoard) GetSize() *[2]int {
	return t.board.GetSize()
}

func (t *TicTacToeBoard) GetBoard() *[][]string {
	return t.board.GetBoard()
}

func (t *TicTacToeBoard) isValidState(state string) bool {
	return (state == t.x || state == t.o || state == t.n)
}

func (t *TicTacToeBoard) isValidPlay(state string) bool {
	return (state == t.x || state == t.o)
}

func (t *TicTacToeBoard) SetLastPlay(play string) {
	t.lastPlay = play
}

func (t *TicTacToeBoard) GetLastPlay() string {
	return t.lastPlay
}

func (t *TicTacToeBoard) SetlastMove(position *[2]int) {
	t.lastMove = *position
}

func (t *TicTacToeBoard) GetlastMove() [2]int {
	return t.lastMove
}

func (t *TicTacToeBoard) SetState(pos *[2]int, state string) (bool, error) {
	if !t.isValidPlay(state) {
		return false, errors.New("invalid state value")
	}
	isStateSet, err := t.board.SetState(pos, state)
	if err != nil {
		return false, err
	}
	if !isStateSet {
		return false, errors.New("failed to set state for tic tac toe board")
	}
	// t.setLastPlay(state)
	// t.setlastMove(pos)
	return true, nil
}

func (t *TicTacToeBoard) GetState(pos *[2]int) (string, error) {
	state, err := t.board.GetState(pos)
	if err != nil {
		return "", err
	}
	if !(t.isValidState(state)) {
		return "", errors.New("invalid state returned")
	}
	return state, nil
}
