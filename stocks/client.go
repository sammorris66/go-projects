package stocks

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// APIClient manages HTTP clients
type APIClient struct {
	Timeout int
}

// NewAPIClient initializes an APIClient
func NewAPIClient() *APIClient {
	return &APIClient{Timeout: 10}
}

// UpdateTimeout updates the client's timeout
func (apiClient *APIClient) UpdateTimeout(timeout int) *APIClient {
	apiClient.Timeout = timeout
	return apiClient
}

// CreateClient returns an HTTP client
func (apiClient *APIClient) CreateClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(apiClient.Timeout) * time.Second,
	}
}

func (apiClient *APIClient) GetRequest(url string, token string) ([]byte, error) {

	// TODO: Add error handling for HTTPStatus OK
	client := apiClient.CreateClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	return body, nil
}
