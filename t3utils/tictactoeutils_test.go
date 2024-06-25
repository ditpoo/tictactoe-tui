package t3utils

import (
	"fmt"
	"testing"
)

func printBoard(board *[][]string) string {
	str := "\n"
	for _, rw := range *board {
		for j, el := range rw {
			if j > 0 {
				str += fmt.Sprintf("| %+v ", el)
			} else {
				str += fmt.Sprintf("%+v ", el)
			}
		}
		str += "\n"
	}
	return str
}

// Test CheckIfWin
func TestCheckIfWin(t *testing.T) {
	// the folowing data is generated using AI (chatgpt4)
	winConditions := [][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}
	for _, cond := range winConditions {
		board := make([][]string, 3)
		for i := range board {
			board[i] = make([]string, 3)
		}
		for r := range board {
			for c := range board {
				board[r][c] = " "
			}
		}
		for _, pos := range cond {
			r, c := pos[0], pos[1]
			board[r][c] = "X"
		}
		ifwin := CheckIfWin(&board, "X")
		if !ifwin {
			t.Fatalf(`failed for win condition, %v for board as below: %v`, cond, printBoard(&board))
		}
	}
}

type Position struct {
	Player string
	Row    int
	Col    int
}

func TestCheckIfDraw(t *testing.T) {
	// the folowing data is generated using AI (chatgpt4)
	drawPositions := [][][3]Position{
		{
			{{Player: "X", Row: 0, Col: 0}, {Player: "O", Row: 0, Col: 1}, {Player: "X", Row: 0, Col: 2}},
			{{Player: "X", Row: 1, Col: 0}, {Player: "X", Row: 1, Col: 1}, {Player: "O", Row: 1, Col: 2}},
			{{Player: "O", Row: 2, Col: 0}, {Player: "X", Row: 2, Col: 1}, {Player: "O", Row: 2, Col: 2}},
		},
		{
			{{Player: "O", Row: 0, Col: 0}, {Player: "X", Row: 0, Col: 1}, {Player: "O", Row: 0, Col: 2}},
			{{Player: "O", Row: 1, Col: 0}, {Player: "O", Row: 1, Col: 1}, {Player: "X", Row: 1, Col: 2}},
			{{Player: "X", Row: 2, Col: 0}, {Player: "O", Row: 2, Col: 1}, {Player: "X", Row: 2, Col: 2}},
		},
		{
			{{Player: "X", Row: 0, Col: 0}, {Player: "X", Row: 0, Col: 1}, {Player: "O", Row: 0, Col: 2}},
			{{Player: "O", Row: 1, Col: 0}, {Player: "O", Row: 1, Col: 1}, {Player: "X", Row: 1, Col: 2}},
			{{Player: "X", Row: 2, Col: 0}, {Player: "X", Row: 2, Col: 1}, {Player: "O", Row: 2, Col: 2}},
		},
		{
			{{Player: "X", Row: 0, Col: 0}, {Player: "O", Row: 0, Col: 1}, {Player: "X", Row: 0, Col: 2}},
			{{Player: "O", Row: 1, Col: 0}, {Player: "O", Row: 1, Col: 1}, {Player: "X", Row: 1, Col: 2}},
			{{Player: "X", Row: 2, Col: 0}, {Player: "X", Row: 2, Col: 1}, {Player: "O", Row: 2, Col: 2}},
		},
	}
	for _, cond := range drawPositions {
		board := make([][]string, 3)
		for i := range board {
			board[i] = make([]string, 3)
		}
		for r := range board {
			for c := range board {
				board[r][c] = " "
			}
		}
		// var rm, cm int
		// var ply string
		for _, Pos := range cond {
			for _, pos := range Pos {
				r, c := pos.Row, pos.Col
				board[r][c] = pos.Player
				// rm, cm = r, c
				// ply = pos.Player
			}
		}
		ifdraw := CheckIfDraw(&board)
		if !ifdraw {
			t.Fatalf(`failed for draw condition, %v for board as below: %v`, cond, printBoard(&board))
		}
	}
}

func TestHasGameEnded(t *testing.T) {
	// the folowing data is generated using AI (chatgpt4)
	winConditions := [][3][2]int{
		{{0, 0}, {0, 1}, {0, 2}},
		{{1, 0}, {1, 1}, {1, 2}},
		{{2, 0}, {2, 1}, {2, 2}},
		{{0, 0}, {1, 0}, {2, 0}},
		{{0, 1}, {1, 1}, {2, 1}},
		{{0, 2}, {1, 2}, {2, 2}},
		{{0, 0}, {1, 1}, {2, 2}},
		{{0, 2}, {1, 1}, {2, 0}},
	}
	for _, cond := range winConditions {
		board := make([][]string, 3)
		for i := range board {
			board[i] = make([]string, 3)
		}
		for r := range board {
			for c := range board {
				board[r][c] = " "
			}
		}
		for _, pos := range cond {
			r, c := pos[0], pos[1]
			board[r][c] = "X"
		}
		hasEnded := HasGameEnded(&board)
		if !hasEnded {
			t.Fatalf(`failed for win condition, %v for board as below: %v`, cond, printBoard(&board))
		}
	}

	// the folowing data is generated using AI (chatgpt4)
	drawPositions := [][][3]Position{
		{
			{{Player: "X", Row: 0, Col: 0}, {Player: "O", Row: 0, Col: 1}, {Player: "X", Row: 0, Col: 2}},
			{{Player: "X", Row: 1, Col: 0}, {Player: "X", Row: 1, Col: 1}, {Player: "O", Row: 1, Col: 2}},
			{{Player: "O", Row: 2, Col: 0}, {Player: "X", Row: 2, Col: 1}, {Player: "O", Row: 2, Col: 2}},
		},
		{
			{{Player: "O", Row: 0, Col: 0}, {Player: "X", Row: 0, Col: 1}, {Player: "O", Row: 0, Col: 2}},
			{{Player: "O", Row: 1, Col: 0}, {Player: "O", Row: 1, Col: 1}, {Player: "X", Row: 1, Col: 2}},
			{{Player: "X", Row: 2, Col: 0}, {Player: "O", Row: 2, Col: 1}, {Player: "X", Row: 2, Col: 2}},
		},
		{
			{{Player: "X", Row: 0, Col: 0}, {Player: "X", Row: 0, Col: 1}, {Player: "O", Row: 0, Col: 2}},
			{{Player: "O", Row: 1, Col: 0}, {Player: "O", Row: 1, Col: 1}, {Player: "X", Row: 1, Col: 2}},
			{{Player: "X", Row: 2, Col: 0}, {Player: "X", Row: 2, Col: 1}, {Player: "O", Row: 2, Col: 2}},
		},
		{
			{{Player: "X", Row: 0, Col: 0}, {Player: "O", Row: 0, Col: 1}, {Player: "X", Row: 0, Col: 2}},
			{{Player: "O", Row: 1, Col: 0}, {Player: "O", Row: 1, Col: 1}, {Player: "X", Row: 1, Col: 2}},
			{{Player: "X", Row: 2, Col: 0}, {Player: "X", Row: 2, Col: 1}, {Player: "O", Row: 2, Col: 2}},
		},
		{
			{{Player: "X", Row: 0, Col: 0}, {Player: "O", Row: 0, Col: 1}, {Player: "X", Row: 0, Col: 2}},
			{{Player: "O", Row: 1, Col: 0}, {Player: "X", Row: 1, Col: 1}, {Player: "O", Row: 1, Col: 2}},
			{{Player: "O", Row: 2, Col: 0}, {Player: "X", Row: 2, Col: 1}, {Player: "X", Row: 2, Col: 2}},
		},
	}
	for _, cond := range drawPositions {
		board := make([][]string, 3)
		for i := range board {
			board[i] = make([]string, 3)
		}
		for r := range board {
			for c := range board {
				board[r][c] = " "
			}
		}
		for _, Pos := range cond {
			for _, pos := range Pos {
				r, c := pos.Row, pos.Col
				board[r][c] = pos.Player
			}
		}
		hasEnded := HasGameEnded(&board)
		if !hasEnded {
			t.Fatalf(`failed for draw condition, %v for board as below: %v`, cond, printBoard(&board))
		}
	}
}
