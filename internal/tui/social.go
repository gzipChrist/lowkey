package tui

import "github.com/charmbracelet/lipgloss"

// TODO: Need to think about where this stuff should live, but prob outside of the tui package.
type socialStatus struct {
	social social
	status status
}

type status int

const (
	unavailable status = iota - 1
	pending
	inflight
	available
)

var ps = map[status]string{
	unavailable: lipgloss.NewStyle().SetString("x").Foreground(lipgloss.Color("#ff0040")).Bold(true).String(),
	inflight:    lipgloss.NewStyle().SetString("...").Faint(true).String(),
	available:   lipgloss.NewStyle().SetString("✓").Foreground(lipgloss.Color("#3fe491")).String(),
}

func (s status) String() string {
	return ps[s]
}

type social int

const (
	snapchat social = iota + 1
	instagram
	tiktok
	twitch
	github
	youtube
	facebook
	mastodon
)

var pm = map[social]string{
	snapchat:  "Snapchat",
	instagram: "Instagram",
	tiktok:    "TikTok",
	twitch:    "Twitch",
	github:    "GitHub",
	youtube:   "YouTube",
	facebook:  "Facebook",
	mastodon:  "Mastodon",
}

func (s social) String() string {
	return pm[s]
}

var su = map[social]string{
	snapchat:  "https://www.snapchat.com/en-US/download",
	instagram: "https://www.instagram.com/accounts/emailsignup",
	tiktok:    "https://www.tiktok.com/signup",
	twitch:    "https://www.twitch.tv/signup",
	github:    "https://github.com/signup",
	youtube:   "https://www.youtube.com",
	facebook:  "https://www.facebook.com/r.php",
	mastodon:  "https://mastodon.social/auth/sign_up",
}

func (s social) SignUpLink() string {
	return su[s]
}

var bm = map[social]string{
	snapchat:  "https://www.snapchat.com/add/",
	instagram: "https://instagram.com/",
	tiktok:    "https://us.tiktok.com/@",
	twitch:    "https://www.twitch.tv/",
	github:    "https://github.com/",
	youtube:   "https://youtube.com/@",
	facebook:  "https://www.facebook.com/",
	mastodon:  "https://mastodon.social/@",
}

func (s social) BaseUrl() string {
	return bm[s]
}
