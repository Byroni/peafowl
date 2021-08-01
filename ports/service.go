package ports

type Peafowl interface {
	// SetDirectory Set directory of saved game clips. TODO: Allow setting and reading from multiple directories
	SetDirectory(path string)

	// SetDiscordChannel Set the name of the discord channel
	//SetDiscordChannel(name string)

	// SetUser Set the user that is posting the clip
	SetUser(name string)

	// SetWebhookURL Set the url that is used to post messages
	SetWebhookURL(url string)

	// GetDirectory returns specified directory name
	GetDirectory() string

	// GetWebhookURL returns specified discord channel
	GetWebhookURL() string

	// GetUser returns specified user
	GetUser() string

	// List Read and list all files in directory set with SetDirectory()
	List()

	// UploadAndPublish Will upload to Gfycat then subsequently publish the gif to the channel set with SetDiscordChannel()
	UploadAndPublish()

	// Upload will upload the file to Gfycat
	Upload()

	// Publish will publish an existing uploaded file to the channel set with SetDiscordChannel
	Publish(message string) error
}
