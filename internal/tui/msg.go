package tui

import "github.com/charmbracelet/lipgloss"

type availability bool

func (a availability) Status() platformStatus {
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

type PlatformResult struct {
	platform    platform
	isAvailable availability
}

type InFlightMsg platformStatus
type PendingStatus platformStatus

type ErrMsg error
