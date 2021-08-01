package gfycat

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

const uploadURL = "https://filedrop.gfycat.com"
const baseURL = "https://api.gfycat.com/v1"
const oAuthEndpoint = "/oauth/token"
const fetchUrlEndpoint = "/gfycat"

type service struct {

}

func New() service {
	return service{}
}

// GetToken Returns bearer token. Returns an empty string if there was an error
func (s *service) GetToken() string {
	resp, err := http.Get("")
	if err != nil {
		log.Error(err)
		return ""
	}

	log.Info(resp)
	return "my token"
}

func (s *service) GenerateFetchURL() string {

	// Generate a new fetch url with baseURL+fetchUrlEndpoint
	// save the newly created url in the database in a new record associated with the file we want to upload
	// store or return the file name returned from gfycat
	return ""
}

func (s *service) Upload() error {
	// The GfyCat API requires use to upload the file that matches exactly the provided gfycat name given by the GenerateFetchURL method.
	// The quick and dirty solution would be to copy the file with the new name, upload, then immediately delete the file.
	// The filename, upload date, new file name, and gfycat  url will be saved in the database.
	return nil
}
