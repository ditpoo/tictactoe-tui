package t3utils

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	x = "X"
	o = "O"
	n = " "
)

func Toss() string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	if rand.Intn(2) == 0 {
		return x
	} else {
		return o
	}
}

// assumes that the play is valid state for tic tac toe
func TogglePlay(play string) string {
	if play == x {
		return o
	}
	return x
}

func FirstMove() [2]int {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	return [2]int{rand.Intn(3), rand.Intn(3)}
}

func CopyBoard(board *[][]string) *[][]string {
	c := make([][]string, len(*board))
	for i := range *board {
		c[i] = make([]string, len((*board)[i]))
		copy(c[i], (*board)[i])
	}
	return &c
}

func PrintBoard(board *[][]string) {
	// fmt.Printf("%+v ", *board)
	fmt.Printf("\n")

	for _, rw := range *board {
		for j, el := range rw {
			if j > 0 {
				fmt.Printf("| %+v ", el)
			} else {
				fmt.Printf("%+v ", el)
			}
		}
		fmt.Printf("\n")
	}
}

func GetPossibleMoves(board *[][]string) [][2]int {
	possibleMoves := make([][2]int, 0)
	for r, rw := range *board {
		for c, el := range rw {
			if el == n {
				possibleMoves = append(possibleMoves, [2]int{r, c})
			}
		}
	}
	return possibleMoves
}

func HasGameStarted(board *[][]string) bool {
	possibleMoves := GetPossibleMoves(board)
	return len(possibleMoves) != 9
}

func GetPlayedMoves(board *[][]string) [][2]int {
	possibleMoves := make([][2]int, 0)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if (*board)[r][c] != n {
				possibleMoves = append(possibleMoves, [2]int{r, c})
			}
		}
	}
	return possibleMoves
}

func GetPlayedMovesByPlayer(board *[][]string, play string) [][2]int {
	possibleMoves := make([][2]int, 0)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if (*board)[r][c] != n && (*board)[r][c] == play {
				possibleMoves = append(possibleMoves, [2]int{r, c})
			}
		}
	}
	return possibleMoves
}

func IsInSetOfMoves(moves [][2]int, move [2]int) bool {
	for i := 0; i < len(moves); i++ {
		r, c := move[0], move[1]
		rm, cm := moves[i][0], moves[i][1]

		if r == rm && c == cm {
			return true
		}
	}
	return false
}

// this assumes that game has not ended and borad is valid tic tac toe one
func GetRandomMove(board *[][]string) [2]int {
	possibleMoves := GetPossibleMoves(board)
	firstMove := FirstMove()

	for !IsInSetOfMoves(possibleMoves, firstMove) {
		firstMove = FirstMove()
	}

	return firstMove
}

func CheckIfWin(board *[][]string, play string) bool {
	tboard := *board
	// "W" indicates win

	// check if there is match row wise
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if tboard[r][c] != play {
				break
			}
			if c == 2 {
				return true
			}
		}
	}

	// check if there is match coloum wise
	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			if tboard[r][c] != play {
				break
			}
			if r == 2 {
				return true
			}
		}
	}

	// check if there is match diagonal wise
	for i := 0; i <= 2; i++ {
		if tboard[i][i] != play {
			break
		}
		if i == 2 {
			return true
		}
	}

	for i := 0; i <= 2; i++ {
		if tboard[i][(2-i)] != play {
			break
		}

		if i == 2 {
			return true
		}
	}

	return false
}

func CheckIfWinMove(board *[][]string, play string, move [2]int) bool {
	rm, cm := move[0], move[1]
	tboard := *board

	// check if there is match row wise
	var r, c, ct, i int
	for r, ct = 0, 0; r < 3; r++ {
		if tboard[r][cm] != play {
			break
		} else {
			ct++
		}
	}
	if ct > 2 {
		return true
	}

	// check if there is match coloum wise
	for c, ct = 0, 0; c < 3; c++ {
		if tboard[rm][c] != play {
			break
		} else {
			ct++
		}
	}
	if ct > 2 {
		return true
	}

	// check if there is match diagonal wise

	for i, ct = 0, 0; i <= 2; i++ {
		if tboard[i][i] != play {
			break
		} else {
			ct++
		}
	}
	if ct > 2 {
		return true
	}

	for i, ct = 0, 0; i <= 2; i++ {
		if tboard[i][(2-i)] != play {
			break
		} else {
			ct++
		}
	}
	return ct > 2
}

func CheckIfDraw(board *[][]string) bool {
	// can any player win in remaining moves
	if !HasGameStarted(board) {
		return false
	}
	if CheckIfWin(board, x) || CheckIfWin(board, o) {
		return false
	}
	// no moves left
	possibleMoves := GetPossibleMoves(board)
	return len(possibleMoves) == 0
}

func HasGameEnded(board *[][]string) bool {
	if !HasGameStarted(board) {
		return false
	}

	if CheckIfDraw(board) {
		return true
	}

	if CheckIfWin(board, x) || CheckIfWin(board, o) {
		return true
	}

	return false
}

func CheckIfGameIsWinnable(board *[][]string, lastPlay string, lastMove [2]int) bool {
	return !HasGameEnded(board)
}

// assuming it is called after the game has ended
func GetWinRow(board *[][]string, play string) [3][2]int {
	tboard := *board
	// "W" indicates win

	// check if there is match row wise
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if tboard[r][c] != play {
				break
			}
			if c == 2 {
				return [3][2]int{{r, 0}, {r, 1}, {r, 2}}
			}
		}
	}

	// check if there is match coloum wise
	for c := 0; c < 3; c++ {
		for r := 0; r < 3; r++ {
			if tboard[r][c] != play {
				break
			}
			if r == 2 {
				return [3][2]int{{0, c}, {1, c}, {2, c}}
			}
		}
	}

	// check if there is match diagonal wise
	for i := 0; i <= 2; i++ {
		if tboard[i][i] != play {
			break
		}
		if i == 2 {
			return [3][2]int{{0,0}, {1, 1}, {2, 2}}
		}
	}

	for i := 0; i <= 2; i++ {
		if tboard[i][(2-i)] != play {
			break
		}

		if i == 2 {
			return [3][2]int{{0, 2}, {1, 1}, {2, 0}}
		}
	}

	return [3][2]int{{-1, -1}, {-1, -1}, {-1, -1}}
}
