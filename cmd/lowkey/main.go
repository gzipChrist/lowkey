package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"payton/internal/tui"
)

func main() {
	m := tui.NewModel()

	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
