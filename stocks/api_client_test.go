package stocks

import (
	"net/http"
	"os"
	"testing"
	"time"

	// cSpell:ignore httpmock
	"github.com/jarcoal/httpmock"
)

func TestNewAPIClient(t *testing.T) {
	client, err := NewAPIClient()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Fatalf("Expected a valid APIClient instance, got nil")
	}
	if client.timeout != 10 {
		t.Errorf("Expected timeout=10, got %d", client.timeout)
	}
}

func TestUpdateToken(t *testing.T) {
	testcases := []struct {
		name      string
		envValue  string
		expectErr bool
	}{
		{"Valid token", "valid_token", false},
		{"Missing token", "", true},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("API_TOKEN", tc.envValue)
			defer os.Unsetenv("API_TOKEN")

			apiClient := &APIClient{}
			err := apiClient.updateToken()

			if err == nil && tc.expectErr {
				t.Errorf("There was no error but there should be no error")
			}

			if apiClient.token != tc.envValue {
				t.Errorf("Expected token %s, got %s", tc.envValue, apiClient.token)
			}
		})
	}

}

func TestUpdateTimeout(t *testing.T) {
	testcases := []struct {
		name      string
		timeout   int
		expectErr bool
	}{
		{"Valid timeout", 20, false},
		{"Invalid timeout", -1, true},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			apiClient := &APIClient{}

			err := apiClient.UpdateTimeout(tc.timeout)

			if err == nil && tc.expectErr {
				t.Errorf("There was no error but there should be no error")
			}

			if (apiClient.timeout != tc.timeout) && !tc.expectErr {
				t.Errorf("Expected timeout %d, got %d", tc.timeout, apiClient.timeout)
			}
		})
	}
}

func TestCreateClient(t *testing.T) {
	testCases := []struct {
		name         string
		timeout      int
		expectedTime time.Duration
	}{
		{"Default timeout", 10, 10 * time.Second},
		{"Zero timeout", 0, 0 * time.Second},
		{"High timeout", 60, 60 * time.Second},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := &APIClient{timeout: tc.timeout}
			httpClient := client.CreateClient()

			if httpClient.Timeout != tc.expectedTime {
				t.Errorf("Expected timeout %v, got %v", tc.expectedTime, client.timeout)
			}
		})
	}
}

func TestGetRequest(t *testing.T) {

	testCases := []struct {
		name         string
		responseBody string
		statusCode   int
		expectError  bool
	}{
		{"Success - 200 OK", `{"message": "success"}`, http.StatusOK, false},
		{"Not Found - 404", ``, http.StatusNotFound, true},
		{"Bad Gateway - 502", ``, http.StatusBadGateway, true},
		{"Internal Server Error - 500", ``, http.StatusInternalServerError, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Activate httpmock
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			// Mock API response
			mockURL := "https://api.example.com/data"
			mockResponse := tc.responseBody

			// Register the mock response
			httpmock.RegisterResponder("GET", mockURL,
				httpmock.NewStringResponder(tc.statusCode, mockResponse))

			// Create APIClient with a test token
			apiClient := &APIClient{token: "test_token"}

			// Call GetRequest
			body, err := apiClient.GetRequest(mockURL)
			if tc.expectError && err == nil {
				t.Errorf("Expected an error but got nil")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			// Verify response body
			if string(body) != mockResponse {
				t.Errorf("Expected response body %q, got %q", mockResponse, string(body))
			}

			// Ensure the request was made
			info := httpmock.GetCallCountInfo()
			if info["GET "+mockURL] != 1 {
				t.Errorf("Expected 1 GET request to %s, got %d", mockURL, info["GET "+mockURL])
			}
		})
	}
}
