package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func HeaderView() string {
	header := `
    __           __           
   / /__ _    __/ /_____ __ __
  / / _ \ |/|/ /  '_/ -_) // /
 /_/\___/__,__/_/\_\\__/\_, / 
                       /___/

 Check username availability
`
	return lipgloss.NewStyle().SetString(header).Foreground(lipgloss.Color("#f4d02d")).String()
}

func SocialListView(state []socialStatus) string {
	var s string

	for _, ps := range state {
		var l string
		if ps.status == available {
			l = ps.social.SignUpLink()
		}
		s += fmt.Sprintf("   %s\t%s  %s\n", ps.social.String(), ps.status.String(), l)
	}

	return s
}

func HelpView() string {
	return lipgloss.NewStyle().SetString(" ctrl+c to quit").Faint(true).String()
}
