package peafowl

import (
	"github.com/byroni/peafowl/ports"
	log "github.com/sirupsen/logrus"
	"strings"
)

const appConfigBucket = "app_config"

type peafowl struct {
	discord  ports.ChatService
	database ports.Database
	// gifService
	directory string
	discordChannel string
	webhookURL string
	user string
}

// New Creates a new peafowl service
func New(cs ports.ChatService, db ports.Database) *peafowl {
	peafowl := &peafowl{
		discord:        cs,
		database:       db,
		directory:      "./",
	}

	// Startup configuration
	user := peafowl.GetUser()
	if len(user) > 0 {
		peafowl.SetUser(user)
	}
	url := peafowl.GetWebhookURL()
	if len(url) > 0 {
		peafowl.SetWebhookURL(url)
	}

	return peafowl
}

func (c *peafowl) SetDirectory(path string) {
	c.directory = path
}

func (c *peafowl) SetUser(name string) {
	c.user = name
	err := c.database.Update(appConfigBucket, "user", name)
	if err != nil {
		log.Error(err)
	}
}

func (c *peafowl) SetWebhookURL(url string) {
	c.webhookURL = url
	c.discord.SetWebhookURL(url)
	err := c.database.Update(appConfigBucket, "webhook_url", url)
	if err != nil {
		log.Error(err)
	}
}

func (c *peafowl) GetDirectory() string {
	return c.directory
}

// GetWebhookURL Returns the set webhookURL.
// If none is set, or there was an error, it returns an empty string
func (c *peafowl) GetWebhookURL() string {
	if len(c.webhookURL) > 0 {
		return c.webhookURL
	}
	value, err := c.database.Get(appConfigBucket, "webhook_url")
	if err != nil {
		log.Error(err)
		return ""
	}
	return value
}

// GetUser Returns the set user.
// If none is set, or there was an error, it returns an empty string
func (c *peafowl) GetUser() string {
	if len(c.user) > 0 {
		return c.user
	}
	value, err := c.database.Get(appConfigBucket, "user")
	if err != nil {
		log.Error(err)
		return ""
	}
	return value
}

// List should list all files in a directory by name.
// Any files that have already been uploaded should be tagged as such.
func (c *peafowl) List() {
}

func (c *peafowl) Upload() {
	// Upload to Gfycat
}

func (c *peafowl) Publish(message string) error {
	messageSlice := []string{
		"@",
		c.user,
		"\n",
		message,
	}

	return c.discord.PublishMessage(strings.Join(messageSlice, ""))
}

func (c *peafowl) UploadAndPublish() {

}
