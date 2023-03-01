package tui

import (
	"context"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"net/http"
	"strings"
	"time"
)

// Put tea.Cmds here.
// TODO: Lots of repeated logic, can DRY all these into one method or split in two by socials and websites.
// Also look into pulling this off the MainModel receiver to keep the receiver for Bubble Tea architecture only.

type platform int

const (
	snapchat platform = iota + 1
	instagram
	tiktok
	twitch
	github
	youtube
	facebook
	mastodon
	dotcom
	dotdev
	dotio
	dotsh
)

var pm = map[platform]string{
	snapchat:  "Snapchat",
	instagram: "Instagram",
	tiktok:    "TikTok",
	twitch:    "Twitch",
	github:    "GitHub",
	youtube:   "YouTube",
	facebook:  "Facebook",
	mastodon:  "Mastodon",
	//dotcom:    ".com",
	//dotdev:    ".dev",
	//dotio:     ".io",
	//dotsh:     ".sh",
}

func (p platform) String() string {
	return pm[p]
}

var su = map[platform]string{
	snapchat:  "https://www.snapchat.com/en-US/download",
	instagram: "https://www.instagram.com/accounts/emailsignup",
	tiktok:    "https://www.tiktok.com/signup",
	twitch:    "https://www.twitch.tv/signup",
	github:    "https://github.com/signup",
	youtube:   "https://www.youtube.com",
	facebook:  "https://www.facebook.com/r.php",
	mastodon:  "https://mastodon.social/auth/sign_up",
}

func (p platform) SignUpLink() string {
	return su[p]
}

var bm = map[platform]string{
	snapchat:  "https://www.snapchat.com/add/",
	instagram: "https://instagram.com/",
	tiktok:    "https://us.tiktok.com/@",
	twitch:    "https://www.twitch.tv/",
	github:    "https://github.com/",
	youtube:   "https://youtube.com/@",
	facebook:  "https://www.facebook.com/",
	mastodon:  "https://mastodon.social/@",
}

func (p platform) BaseUrl() string {
	return bm[p]
}

func setStatusPending() tea.Msg {
	return PendingStatus(pending)
}

func setStatusInFlight() tea.Msg {
	return InFlightMsg(inflight)
}

func checkSnapchat(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := snapchat.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ErrMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ErrMsg(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrMsg(err)
		}

		if strings.Contains(string(body), "content=\"Not_Found\"") {
			return PlatformResult{
				platform:    snapchat,
				isAvailable: true,
			}
		}

		return PlatformResult{
			platform:    snapchat,
			isAvailable: false,
		}
	}

}

func checkInstagram(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := instagram.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ErrMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ErrMsg(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrMsg(err)
		}

		if strings.Contains(string(body), "<title>Instagram</title>") {
			return PlatformResult{
				platform:    instagram,
				isAvailable: true,
			}
		}

		return PlatformResult{
			platform:    instagram,
			isAvailable: false,
		}
	}
}

func checkTikTok(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := tiktok.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ErrMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ErrMsg(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrMsg(err)
		}

		if strings.Contains(string(body), "Watch the latest video from .") {
			return PlatformResult{
				platform:    tiktok,
				isAvailable: true,
			}
		}

		return PlatformResult{
			platform:    tiktok,
			isAvailable: false,
		}
	}
}

func checkTwitch(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := twitch.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ErrMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ErrMsg(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrMsg(err)
		}

		if strings.Contains(string(body), "content='Twitch is the world") {
			return PlatformResult{
				platform:    twitch,
				isAvailable: true,
			}
		}

		return PlatformResult{
			platform:    twitch,
			isAvailable: false,
		}
	}
}

func checkGitHub(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := github.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return PlatformResult{
				platform:    github,
				isAvailable: true,
			}
		}

		return PlatformResult{
			platform:    github,
			isAvailable: false,
		}
	}
}

func checkYouTube(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := youtube.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode == http.StatusNotFound || strings.Contains(string(body), "<title>404 Not Found</title>") {
			return PlatformResult{
				platform:    youtube,
				isAvailable: true,
			}
		}

		return PlatformResult{
			platform:    youtube,
			isAvailable: false,
		}
	}
}

func checkFacebook(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := facebook.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ErrMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ErrMsg(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrMsg(err)
		}

		if strings.Contains(string(body), "<title>") {
			return PlatformResult{
				platform:    facebook,
				isAvailable: false,
			}
		}

		return PlatformResult{
			platform:    facebook,
			isAvailable: true,
		}
	}
}

func checkMastodon(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		url := mastodon.BaseUrl() + username

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return ErrMsg(err)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ErrMsg(err)
		}

		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return ErrMsg(err)
		}

		if strings.Contains(string(body), "<title>The page you are looking for") || strings.Contains(string(body), "<title>The page you were looking for") {
			return PlatformResult{
				platform:    mastodon,
				isAvailable: true,
			}
		}

		return PlatformResult{
			platform:    mastodon,
			isAvailable: false,
		}
	}
}

// TODO: Implement website checking logic.
//func (m MainModel) checkDotCom() tea.Msg {
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//
//	d := dotcom
//	url := fmt.Sprintf("https://www.%s%s", m.username, d.String())
//
//	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
//	if err != nil {
//		return err
//	}
//
//	resp, err := http.DefaultClient.Do(req)
//	if err == context.DeadlineExceeded || resp == nil || resp.StatusCode == http.StatusNotFound {
//		return PlatformResult{
//			platform:    dotcom,
//			isAvailable: true,
//		}
//	}
//
//	return PlatformResult{
//		platform:    dotcom,
//		isAvailable: false,
//	}
//
//}
