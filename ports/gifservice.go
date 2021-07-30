package ports

type GifService interface {
	// Upload Uploads file to our gif service
	Upload() bool
}
