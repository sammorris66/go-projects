package stocks

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// APIClient manages HTTP clients
type APIClient struct {
	timeout int
	token   string
}

// NewAPIClient initializes an APIClient
func NewAPIClient() (*APIClient, error) {
	apiClient := &APIClient{
		timeout: 10,
	}
	err := apiClient.updateToken()
	if err != nil {
		return nil, err
	}
	return apiClient, nil
}

func (apiClient *APIClient) updateToken() error {

	token := os.Getenv("API_TOKEN")

	if token == "" {
		return fmt.Errorf("api token variable is blank")
	}

	apiClient.token = token
	return nil

}

// UpdateTimeout updates the client's timeout
func (apiClient *APIClient) UpdateTimeout(timeout int) error {
	if timeout <= 0 {
		return fmt.Errorf("timeout must be greater than zero, got %d", timeout)
	}

	apiClient.timeout = timeout
	return nil
}

// CreateClient returns an HTTP client
func (apiClient *APIClient) CreateClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(apiClient.timeout) * time.Second,
	}
}

func (apiClient *APIClient) GetRequest(url string) ([]byte, error) {

	client := apiClient.CreateClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Add("Authorization", "Bearer "+apiClient.token)

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
