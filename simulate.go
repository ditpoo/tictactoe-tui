package main

import (
	"fmt"
	"t3/board"
	"t3/t3board"
	"t3/t3gai"
	"t3/t3game"
	"t3/t3rules"
	"t3/t3utils"
)

func DebugGameMove(T3G *t3game.TicTacToeGame, play string, move [2]int) {
	check, err := T3G.SetMove(play, &move)
	if err != nil {
		fmt.Println(check, err)
		return
	}
	t3utils.PrintBoard(T3G.GetBoard())
}

func NewStandardTicTacToeGame() (*t3game.TicTacToeGame, error) {
	B, err := board.NewSquare2DBoard(3, 3)
	if err != nil {
		return nil, err
	}
	T3B, err := t3board.NewTicTacToeBoard(B)
	if err != nil {
		return nil, err
	}
	T3R := t3rules.NewStandardTicTacToeGameRules()
	T3G, err := t3game.NewTicTacToeGame(*T3B, T3R)
	if err != nil {
		return nil, err
	}
	return T3G, nil
}

func Play() {
	fmt.Println("In Main")

	T3G, err := NewStandardTicTacToeGame()
	if err != nil {
		fmt.Println(err)
		return
	}
	t3utils.PrintBoard(T3G.GetBoard())

	play := t3utils.Toss()
	move := t3gai.CalculateMove(T3G.GetBoard(), play)
	if move == nil {
		fmt.Println("issues with gai")
		t3utils.PrintBoard(T3G.GetBoard())
		return
	}
	DebugGameMove(T3G, play, *move)

	for !T3G.HasGameEnded() {
		play = t3utils.TogglePlay(play)
		move := t3gai.CalculateMove(T3G.GetBoard(), play)
		fmt.Println("player: ", play, "move: ", move)
		if move == nil {
			fmt.Println("issues with gai")
			t3utils.PrintBoard(T3G.GetBoard())
			return
		}
		DebugGameMove(T3G, play, *move)
	}
	fmt.Println(T3G.GetResult())
}
