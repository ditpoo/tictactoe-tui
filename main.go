package main

// Code in this file is made using AI (chatgpt4)
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
	game       *t3game.TicTacToeGame
	player     string
	winner     string
	cursorX    int
	cursorY    int
	aiMove     [2]int // to track AI's last move
	showAiMove bool   // flag to highlight AI's move
	winRow     [3][2]int
}

func main() {
	game, err := NewStandardTicTacToeGame()
	if err != nil {
		fmt.Println(err)
		return
	}
	p := tea.NewProgram(initialModel(game))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %s\n", err)
		os.Exit(1)
	}
}

func initialModel(game *t3game.TicTacToeGame) model {
	return model{
		game:    game,
		player:  game.Toss(),
		cursorX: 0,
		cursorY: 0,
		winRow: [3][2]int{{-1,-1}, {-1,-1}, {-1,-1}},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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
						return m, tea.Tick(time.Second*2, func(time.Time) tea.Msg {
							return "clearHighlight"
						})
					}
				} else {
					m.player = m.game.TogglePlay()
					m.aiMove = *t3gai.CalculateMove(m.game.GetBoard(), m.player)
					m.game.SetMove(m.player, &m.aiMove)
					m.showAiMove = true // Enable highlight for AI move
					m.player = m.game.TogglePlay()
					return m, tea.Tick(time.Second*2, func(time.Time) tea.Msg {
						return "clearHighlight"
					})
				}
			}
		}
	case string: // Handling the clear highlight message
		if msg == "clearHighlight" {
			m.showAiMove = false
			return m, tea.Cmd(nil) // Update the view without any further commands
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString("\n")
	for i, row := range *m.game.GetBoard() {
		for j, cell := range row {
			cellStr := ifEmpty(cell, " ")
			position := [2]int{i, j}
			if m.cursorY == i && m.cursorX == j {
				if inWinRow(position, m) || (m.showAiMove && m.aiMove[0] == i && m.aiMove[1] == j) {
					// Highlight the entire cell for AI move
					b.WriteString(fmt.Sprintf("\033[48;2;240;240;240m\033[38;2;0;0;0m[%s]\033[0m", cellStr))
				} else {
					// Normal cursor highlighting
					b.WriteString(fmt.Sprintf("[%s]", cellStr))
				}
			} else {
				if inWinRow(position, m) || (m.showAiMove && m.aiMove[0] == i && m.aiMove[1] == j) {
					// Highlight the entire cell for AI move
					b.WriteString(fmt.Sprintf("\033[48;2;240;240;240m\033[38;2;0;0;0m %s \033[0m", cellStr))
				} else {
					// Normal cell
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