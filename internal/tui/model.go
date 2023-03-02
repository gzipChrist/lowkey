package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// MainModel is the root state of the app.
type MainModel struct {
	username  string
	textInput textinput.Model
	statuses  []socialStatus
	err       error
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
			if m.textInput.Value() != "" {
				m.username = m.textInput.Value()

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

	case UsernameResultMsg:
		m.statuses[msg.social-1] = socialStatus{
			social: msg.social,
			status: msg.isAvailable.Status(),
		}

	case InFlightMsg:
		for i, p := range m.statuses {
			m.statuses[i] = socialStatus{
				social: p.social,
				status: inflight,
			}
		}

	// TODO: Pass social to err message and print error info.
	case ErrMsg:

	}

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// View renders a string representation of the MainModel.
func (m MainModel) View() string {
	return fmt.Sprintf(
		"%s%s%s%s%s%s%s",
		HeaderView(),
		"\n",
		m.textInput.View(),
		"\n\n",
		SocialListView(m.statuses),
		"\n\n",
		HelpView(),
	)
}

// NewModel configures the initial model at runtime.
func NewModel() MainModel {
	ti := textinput.New()
	ti.Prompt = " @ "
	ti.Placeholder = "Enter username"
	ti.Focus()

	return MainModel{
		textInput: ti,
		statuses: []socialStatus{
			{
				social: snapchat,
				status: pending,
			},
			{
				social: instagram,
				status: pending,
			},
			{
				social: tiktok,
				status: pending,
			},
			{
				social: twitch,
				status: pending,
			},
			{
				social: github,
				status: pending,
			},
			{
				social: youtube,
				status: pending,
			},
			{
				social: facebook,
				status: pending,
			},
			{
				social: mastodon,
				status: pending,
			},
		},
	}
}
