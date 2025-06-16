package telegram

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	Token     string
	ChannelID string
}

func NewClient(token, channelID string) *Client {
	return &Client{
		Token:     token,
		ChannelID: channelID,
	}
}

func (c *Client) SendMessage(message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.Token)

	resp, err := http.PostForm(apiURL, url.Values{
		"chat_id":    {c.ChannelID},
		"text":       {message},
		"parse_mode": {"Markdown"},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram returned status %d: %s", resp.StatusCode, body)
	}
	return nil
}
