package main

import (
    "fmt"
    "os"
    "strings"
	"t3/t3gai"
	"t3/t3game"
	"time"
    tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    game              *t3game.TicTacToeGame
    player            string
    winner            string
    cursorX           int
    cursorY           int
    aiMove            [2]int
    showAiMove        bool
    winRow            [3][2]int
    menu              []string
    selected          int
    inGame            bool
    showInstructions  bool
    instructions      string
	instructionsMap   map[int]string
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Alas, there's been an error: %s\n", err)
        os.Exit(1)
    }
}

func initialModel() model {
    return model{
        menu:             []string{"Standard Tic Tac Toe", "Inverse Tic Tac Toe", "Neutral Tic Tac Toe"},
		instructionsMap:  map[int]string{
			0: standardt3instructions,
			1: inverset3instructions,
			2: neutralt3instructions,
		},
        selected:         0,
        inGame:           false,
        showInstructions: false,
        instructions:     "",
		winRow: [3][2]int{{-1,-1}, {-1,-1}, {-1,-1}},
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.Type == tea.KeyCtrlC || msg.String() == "esc" {
            return m, tea.Quit
        }
        if m.inGame {
            return m.gameUpdate(msg)
        } else if m.showInstructions {
            if msg.Type == tea.KeyEnter {
                m.inGame = true
                m.showInstructions = false
            }
            return m, nil
        }
        switch msg.Type {
        case tea.KeyUp, tea.KeyDown:
            if msg.Type == tea.KeyUp && m.selected > 0 {
                m.selected--
            } else if msg.Type == tea.KeyDown && m.selected < len(m.menu)-1 {
                m.selected++
            }
        case tea.KeyEnter:
            var err error
            m.game, err = initializeGame(m.menu[m.selected])
            if err != nil {
                fmt.Println(err)
                return m, tea.Quit
            }
			m.player = m.game.Toss()
            m.instructions = fmt.Sprintf("Instructions for %s:\n\n%s\n\n", m.menu[m.selected], m.instructionsMap[m.selected])
            m.showInstructions = true
        }
	case string:
		if m.inGame {
            return m.gameUpdate(msg)
        }
    }
    return m, nil
}

const (
	standardt3instructions = `Objective: Align three of your marks (X or O) in a row (horizontally, vertically, or diagonally).

Instructions:
- Players alternate turns placing their mark in an empty square.
- The first to align three marks in a row wins.
- If all squares fill up without a three-in-a-row, it's a draw.

Press Enter to start playing Standard Tic Tac Toe.`
	inverset3instructions = `Objective: Avoid aligning three of your marks in a row. Doing so results in a loss.

Instructions:
- Players alternate turns placing their mark in an empty square.
- If a player lines up three of their marks (in any direction), they lose.
- The game ends in a draw if all squares are filled without any three-in-a-row.

Press Enter to start playing Inverse Tic Tac Toe.
`
	neutralt3instructions = `Objective: Align three marks X in a row, with players not restricted to one mark type.

Instructions:
- Players alternate turns and may place X in an empty square each turn.
- The goal is to line up three identical marks X in a row.

Press Enter to start playing Neutral Tic Tac Toe.
`
)

func initializeGame(choice string) (*t3game.TicTacToeGame, error) {
    switch choice {
    case "Standard Tic Tac Toe":
        return NewStandardTicTacToeGame()
    case "Inverse Tic Tac Toe":
        return NewInverseTicTacToeGame()
    case "Neutral Tic Tac Toe":
        return NewNeutralTicTacToeGame()
    }
    return nil, fmt.Errorf("invalid game choice")
}

func (m model) gameUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
    // Add your existing game logic here...
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "winupdate":
			return m, tea.Cmd(nil)
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "left":
			if m.cursorX > 0 {
				m.cursorX--
			}
		case "right":
			if m.cursorX < 2 {
				m.cursorX++
			}
		case "up":
			if m.cursorY > 0 {
				m.cursorY--
			}
		case "down":
			if m.cursorY < 2 {
				m.cursorY++
			}
		case "enter":
			state, err := m.game.GetState(&[2]int{m.cursorY, m.cursorX})
			if err != nil {
				break
			}
			if m.winner == "" && state == " " {
				isSet, err := m.game.SetMove(m.player, &[2]int{m.cursorY, m.cursorX})
				if err != nil {
					fmt.Println(isSet, err)
					break
				}
				if m.game.HasGameEnded() {
					res := m.game.GetResult()
					if res.IsDraw {
						m.winner = "D"
					} else {
						m.winner = res.Winner
						m.winRow = m.game.GetWinRow()
						return m, tea.Tick(time.Second*1, func(time.Time) tea.Msg {
							return "clearHighlight"
						})
					}
				} else {
					m.player = m.game.TogglePlay()
					m.aiMove = *t3gai.CalculateMove(m.game.GetBoard(), m.player)
					m.game.SetMove(m.player, &m.aiMove)
					m.showAiMove = true
					m.player = m.game.TogglePlay()
					return m, tea.Tick(time.Second*1, func(time.Time) tea.Msg {
						return "clearHighlight"
					})
				}
			}
		}
	case string: // Handling the clear highlight message
		if !m.game.HasGameEnded() && msg == "clearHighlight" {
			m.showAiMove = false
			return m, tea.Cmd(nil) // Update the view without any further commands
		} else if m.game.HasGameEnded() {
			m.showAiMove = false
			res := m.game.GetResult()
			if res.IsDraw {
				m.winner = "D"
				return m, nil
			} else {
				m.winner = res.Winner
				m.winRow = m.game.GetWinRow()
				return m, tea.Tick(time.Second*2, func(time.Time) tea.Msg {
					return "winupdate"
				})
			}
		}
	}
    return m, nil
}

func (m model) View() string {
    if m.showInstructions {
        return m.instructions
    }
    if m.inGame {
        var b strings.Builder
        b.WriteString("\n")
        for i, row := range *m.game.GetBoard() {
            for j, cell := range row {
                cellStr := ifEmpty(cell, " ")
                position := [2]int{i, j}
                if m.cursorY == i && m.cursorX == j {
                    if inWinRow(position, m) || (m.showAiMove && m.aiMove[0] == i && m.aiMove[1] == j) {
                        b.WriteString(fmt.Sprintf("\033[48;2;240;240;240m\033[38;2;0;0;0m[%s]\033[0m", cellStr))
                    } else {
                        b.WriteString(fmt.Sprintf("[%s]", cellStr))
                    }
                } else {
                    if inWinRow(position, m) || (m.showAiMove && m.aiMove[0] == i && m.aiMove[1] == j) {
                        b.WriteString(fmt.Sprintf("\033[48;2;240;240;240m\033[38;2;0;0;0m %s \033[0m", cellStr))
                    } else {
                        b.WriteString(fmt.Sprintf(" %s ", cellStr))
                    }
                }
                if j < 2 {
                    b.WriteString("|")
                }
            }
            if i < 2 {
                b.WriteString("\n---+---+---\n")
            }
        }
        b.WriteString("\n\n")
        if m.winner == "D" {
            b.WriteString("Game is a draw!\n")
        } else if m.winner != "" {
            b.WriteString(fmt.Sprintf("Player %s wins!\n", m.winner))
        } else {
            b.WriteString(fmt.Sprintf("Current Player: %s\n", m.player))
        }
        b.WriteString("Press Esc or Ctrl+C to quit.\n")
        return b.String()
    } else {
		// Menu rendering
		s := "Select a game variant:\n\n"
		for i, choice := range m.menu {
			cursor := " " // no cursor
			if i == m.selected {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}
		return s
	}
}

func ifEmpty(val string, defaultVal string) string {
	if val == "" {
		return defaultVal
	}
	return val
}

func inWinRow(position [2]int, m model) bool {
	rp, cp := position[0], position[1]
	for _, pos:= range m.winRow {
		r, c := pos[0], pos[1]
		if r == rp && c == cp {
			return true
		}
	}
	return false
}