package ports

type ChatService interface {
	SetWebhookURL(url string)
	PublishMessage(string) error
}
