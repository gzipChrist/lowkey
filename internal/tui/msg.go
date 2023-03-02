package tui

import "github.com/charmbracelet/lipgloss"

type availability bool

func (a availability) Status() status {
	s := unavailable
	if a {
		s = available
	}

	return s
}

func (a availability) String() string {
	s := lipgloss.NewStyle().SetString("𐄂").Foreground(lipgloss.Color("#FF004F")).String()
	if a {
		s = lipgloss.NewStyle().SetString("✓").Foreground(lipgloss.Color("#00ffb0")).String()
	}

	return s
}

type UsernameResultMsg struct {
	social      social
	isAvailable availability
}

type InFlightMsg status
type PendingStatus status

type ErrMsg error
