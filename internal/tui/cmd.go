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

func setStatusPending() tea.Msg {
	return PendingStatus(pending)
}

func setStatusInFlight() tea.Msg {
	return InFlightMsg(inflight)
}

func checkSnapchat(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      snapchat,
				isAvailable: true,
			}
		}

		return UsernameResultMsg{
			social:      snapchat,
			isAvailable: false,
		}
	}

}

func checkInstagram(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      instagram,
				isAvailable: true,
			}
		}

		return UsernameResultMsg{
			social:      instagram,
			isAvailable: false,
		}
	}
}

func checkTikTok(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      tiktok,
				isAvailable: true,
			}
		}

		return UsernameResultMsg{
			social:      tiktok,
			isAvailable: false,
		}
	}
}

func checkTwitch(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      twitch,
				isAvailable: true,
			}
		}

		return UsernameResultMsg{
			social:      twitch,
			isAvailable: false,
		}
	}
}

func checkGitHub(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      github,
				isAvailable: true,
			}
		}

		return UsernameResultMsg{
			social:      github,
			isAvailable: false,
		}
	}
}

func checkYouTube(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      youtube,
				isAvailable: true,
			}
		}

		return UsernameResultMsg{
			social:      youtube,
			isAvailable: false,
		}
	}
}

func checkFacebook(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      facebook,
				isAvailable: false,
			}
		}

		return UsernameResultMsg{
			social:      facebook,
			isAvailable: true,
		}
	}
}

func checkMastodon(username string) func() tea.Msg {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
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
			return UsernameResultMsg{
				social:      mastodon,
				isAvailable: true,
			}
		}

		return UsernameResultMsg{
			social:      mastodon,
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
