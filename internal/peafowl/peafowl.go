package peafowl

import (
	"github.com/byroni/peafowl/ports"
	"strings"
)

type peafowl struct {
	discord  ports.ChatService
	database ports.Database
	// gifService
	directory string
	discordChannel string
	user string
}

// New Creates a new peafowl service
func New(cs ports.ChatService, db ports.Database) *peafowl {
	return &peafowl{
		discord:        cs,
		database:       db,
		directory:      "./",
		discordChannel: "clips",
		user:           "Me",
	}
}

func (c *peafowl) SetDirectory(path string) {
	c.directory = path
}
func (c *peafowl) SetDiscordChannel(name string) {
	c.discordChannel = name
}

func (c *peafowl) SetUser(name string) {
	c.user = name
}

func (c *peafowl) GetDirectory() string {
	return c.directory
}

func (c *peafowl) GetDiscordChannel() string {
	return c.discordChannel
}

func (c *peafowl) GetUser() string {
	return c.user
}

func (c *peafowl) List() {

}

func (c *peafowl) Upload() {
	// Upload to Gfycat
}

func (c *peafowl) Publish(message string) {
	messageSlice := []string{
		"@",
		c.user,
		"\n",
		message,
	}

	c.discord.PublishMessage(strings.Join(messageSlice, ""))
}

func (c *peafowl) UploadAndPublish() {

}
