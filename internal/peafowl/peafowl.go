package peafowl

type peafowl struct {
	// discordService
	// gifService
	directory string
	discordChannel string
	user string
}

// TODO: Import dependencies here\

// New Creates a new peafowl service
func New() *peafowl {
	return &peafowl{
		directory: "./",
		discordChannel: "clips",
		user:      "Me",
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

}

func (c *peafowl) Publish() {

}

func (c *peafowl) UploadAndPublish() {

}
