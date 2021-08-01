package discord

import (
	"fmt"
	"net/http"
	"net/url"
)

type Service struct {
	webhookURL string
	user string
}

type Config struct {
	WebhookUrl string
	User string
}

// New creates a new empty service
func New() *Service {
	return &Service{}
}

// NewWithConfig creates a service with loaded configuration
func NewWithConfig(config Config) *Service {
	return &Service{
		webhookURL: config.WebhookUrl,
		user: config.User,
	}
}

// SetWebhookURL Sets the URL to use for sending messages.
// This should be used only if the url is not initially loaded in on startup
func (s *Service) SetWebhookURL(url string) {
	s.webhookURL = url
}

func (s *Service) PublishMessage(message string) error {
	resp, err := http.PostForm(s.webhookURL, url.Values{
		"content": {message},
	})
	if err != nil {
		fmt.Println("Something went wrong")
		return err
	}
	defer resp.Body.Close()
	return nil
}
