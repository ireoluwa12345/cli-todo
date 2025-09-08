package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Tui struct {
}

const (
	padding int = 12
)

func New() *Tui {
	return &Tui{}
}

func (t *Tui) Run(m tea.Model) {
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
