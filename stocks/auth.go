package stocks

import "os"

// APIAuth stores authentication tokens
type APIAuth struct {
	token string
}

func NewAPIAuth() *APIAuth {
	return &APIAuth{}
}

// GetToken retrieves API token from environment variables
func (apiAuth *APIAuth) GetToken() string {
	if apiAuth.token == "" {
		apiAuth.token = os.Getenv("API_TOKEN")
	}
	return apiAuth.token
}
