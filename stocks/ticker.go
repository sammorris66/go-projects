package stocks

import (
	"fmt"
	"net/url"
)

// Interface created to help with mocking test data
type Ticker interface {
	GetTickers() ([]string, error)
}

type StockTicker struct {
	market    string
	APIClient APIClient
}

func NewTicker(market string) (*StockTicker, error) {

	if market == "" {
		return nil, fmt.Errorf("market can not be empty")
	}

	apiClient, err := NewAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}

	return &StockTicker{
		market:    market,
		APIClient: *apiClient,
	}, nil
}

// CreateUrl generates the API endpoint URL
func (t *StockTicker) createUrl() (string, error) {

	basePath, err := url.JoinPath(baseUrl, "/v3/reference/tickers")
	if err != nil {
		return "", fmt.Errorf("error joining path: %w", err)
	}

	parsedURL, err := url.Parse(basePath)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}

	queryParams := url.Values{}
	queryParams.Set("active", "true")
	queryParams.Set("limit", "1000")
	queryParams.Set("market", t.market)

	parsedURL.RawQuery = queryParams.Encode()
	return parsedURL.String(), nil
}

func (t *StockTicker) GetTickers() ([]string, error) {

	var listTickers []string

	url, err := t.createUrl()
	if err != nil {
		return nil, fmt.Errorf("can't create a URL: %w", err)
	}

	resp, err := t.APIClient.GetRequest(url)
	if err != nil {
		return nil, fmt.Errorf("the is a error calling the Get method %w", err)
	}

	tickers, err := t.parseResponseBody(resp)
	if err != nil {
		return nil, fmt.Errorf("there is an error parsing the response %w", err)
	}

	for _, ticker := range tickers.Results {

		listTickers = append(listTickers, ticker.Ticker)
	}

	return listTickers, nil
}

// parseResponseBody parses JSON response
func (StockTicker) parseResponseBody(body []byte) (TickerInformation, error) {
	fxParser := JsonParser[TickerInformation]{}
	parsedBody, err := fxParser.ParseResponseBody(body)
	if err != nil {
		return TickerInformation{}, fmt.Errorf("error parsing response body: %w", err)
	}

	return *parsedBody, nil
}
