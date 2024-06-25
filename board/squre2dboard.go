package board

import (
	"errors"
	"fmt"
)

// Square 2D Board
type Square2DBoard struct {
	rowCount    int
	columnCount int
	board       [][]string
}

type BoardHandler interface {
	GetSize() *[2]int
	GetState(pos *[2]int) (string, error)
	SetState(pos *[2]int, state string) (bool, error)
	GetBoard() *[][]string
}

func NewSquare2DBoard(rowCount int, columnCount int) (*Square2DBoard, error) {
	if !(rowCount > 0) {
		return nil, errors.New("row count is invalid or not set")
	}
	if !(columnCount > 0) {
		return nil, errors.New("row count is invalid or not set")
	}
	slice := make([][]string, rowCount)
	for i := range slice {
		slice[i] = make([]string, columnCount)
	}
	return &Square2DBoard{
		rowCount:    rowCount,
		columnCount: columnCount,
		board:       slice,
	}, nil
}

func (s *Square2DBoard) getRowCount() int {
	return s.rowCount
}

func (s *Square2DBoard) getColumnCount() int {
	return s.columnCount
}

func (s *Square2DBoard) validatePosition(pos *[2]int) (bool, error) {
	rp, cp := pos[0], pos[1]
	if !(rp >= 0 && rp <= s.rowCount) {
		return false, errors.New("invalid row co-ordinate")
	}
	if !(cp >= 0 && cp <= s.columnCount) {
		return false, errors.New("invalid coloumn co-ordinate")
	}

	return true, nil
}

func (s *Square2DBoard) getState(pos *[2]int) (string, error) {
	isValid, err := s.validatePosition(pos)
	if !isValid {
		return "", err
	}

	return s.board[pos[0]][pos[1]], nil
}

func (s *Square2DBoard) setState(pos *[2]int, state string) (bool, error) {
	isValid, err := s.validatePosition(pos)
	if !isValid {
		return false, err
	}
	r, c := pos[0], pos[1]
	fmt.Printf("%+v", s.board[r][c])
	s.board[r][c] = state

	return true, nil
}

func (s *Square2DBoard) getBoard() *[][]string {
	c := make([][]string, len(s.board))
	for i := range s.board {
		c[i] = make([]string, len(s.board[i]))
		copy(c[i], s.board[i])
	}
	return &c
}

func (s *Square2DBoard) GetSize() *[2]int {
	r, c := s.getRowCount(), s.getColumnCount()
	size := [2]int{r, c}
	return &size
}

func (s *Square2DBoard) GetState(pos *[2]int) (string, error) {
	state, err := s.getState(pos)
	if err != nil {
		return "", err
	}

	return state, nil
}

func (s *Square2DBoard) SetState(pos *[2]int, state string) (bool, error) {
	isSet, err := s.setState(pos, state)
	if err != nil {
		return false, err
	}
	if !isSet {
		return false, errors.New("failed to set the state")
	}
	return true, nil
}

func (s *Square2DBoard) GetBoard() *[][]string {
	return s.getBoard()
}
