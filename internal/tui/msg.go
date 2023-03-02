package tui

type UsernameResultMsg struct {
	social      social
	isAvailable availability
}

type InFlightMsg status
type PendingStatus status

type ErrMsg error
