package discord

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
)

type Service struct {
	webhookURL string
}


// New creates a new empty service
func New() *Service {
	return &Service{}
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
		log.Error(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}
