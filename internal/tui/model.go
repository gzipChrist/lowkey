package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status struct {
	platform platform
	status   platformStatus
}

var platformState = []status{
	{
		platform: snapchat,
		status:   pending,
	},
	{
		platform: instagram,
		status:   pending,
	},
	{
		platform: tiktok,
		status:   pending,
	},
	{
		platform: twitch,
		status:   pending,
	},
	{
		platform: github,
		status:   pending,
	},
	{
		platform: youtube,
		status:   pending,
	},
	{
		platform: facebook,
		status:   pending,
	},
	{
		platform: mastodon,
		status:   pending,
	},
}

type platformStatus int

const (
	unavailable platformStatus = iota - 1
	pending
	inflight
	available
)

var ps = map[platformStatus]string{
	unavailable: lipgloss.NewStyle().SetString("x").Foreground(lipgloss.Color("#ff0040")).Bold(true).String(),
	inflight:    lipgloss.NewStyle().SetString("...").Faint(true).String(),
	available:   lipgloss.NewStyle().SetString("✓").Foreground(lipgloss.Color("#3fe491")).String(),
}

func (p platformStatus) String() string {
	return ps[p]
}

// MainModel is the root state of the app.
type MainModel struct {
	username string

	TextInput textinput.Model

	platformState []status

	err error
}

// NewModel configures the initial model at runtime.
func NewModel() MainModel {
	ti := textinput.New()
	ti.Prompt = " @ "
	ti.Placeholder = "Enter username"
	ti.Focus()

	return MainModel{
		TextInput:     ti,
		platformState: platformState,
	}
}

// Init returns any number of tea.Cmds at runtime.
func (m MainModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles all tea.Msgs in the Bubble Tea event loop.
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// Handle keypress messages.
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		// Kick off username search.
		case tea.KeyEnter:
			if m.TextInput.Value() != "" {
				m.username = m.TextInput.Value()

				cmds = append(cmds,
					setStatusInFlight,
					checkSnapchat(m.username),
					checkInstagram(m.username),
					checkTikTok(m.username),
					checkTwitch(m.username),
					checkGitHub(m.username),
					checkYouTube(m.username),
					checkFacebook(m.username),
					checkMastodon(m.username),
				)
			}

		}

	case PlatformResult:
		platformState[msg.platform-1] = status{
			platform: msg.platform,
			status:   msg.isAvailable.Status(),
		}

	case InFlightMsg:
		for i, p := range m.platformState {
			m.platformState[i] = status{
				platform: p.platform,
				status:   inflight,
			}
		}

	// TODO: Pass platform to err message and print error info.
	case ErrMsg:

	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders a string representation of the MainModel.
func (m MainModel) View() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		HeaderView(),
		"\n",
		m.TextInput.View(),
		"\n\n",
		PlatformsListView(m.platformState),
		"\n\n",
		HelpView(),
	)
}
