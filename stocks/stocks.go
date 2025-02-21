package stocks

import (
	"fmt"
	"net/url"
	"slices"
)

const baseUrl string = "https://api.polygon.io"

type Stocks struct {
	symbol    string
	price     float64
	apiClient APIClient
	ticker    Ticker
}

// NewStocks initializes an `Stocks` instance
func NewStocks(symbol, market string) (*Stocks, error) {

	if symbol == "" {
		return nil, fmt.Errorf("symbol can not be empty")
	}

	if market == "" {
		return nil, fmt.Errorf("market can not be empty")
	}

	ticker, err := NewTicker(market) // Capture both values
	if err != nil {
		return nil, fmt.Errorf("failed to create ticker: %w", err)
	}

	apiClient, err := NewAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create api client: %w", err)
	}

	stock := &Stocks{
		symbol:    symbol,
		apiClient: *apiClient,
		ticker:    *ticker,
	}

	valid, err := stock.validateSymbol(symbol)
	if err != nil {
		return nil, fmt.Errorf("symbol validation failed: %w", err)
	}
	if !valid {
		return nil, fmt.Errorf("invalid symbol: %s", symbol)
	}

	return stock, nil

}

func (s *Stocks) validateSymbol(symbol string) (bool, error) {

	resp, err := s.ticker.GetTickers()
	if err != nil {
		return false, fmt.Errorf("can not get the ticker information: %w", err)
	}

	if slices.Contains(resp, symbol) {
		fmt.Printf("Symbol %v is found ", symbol)
		return true, nil
	} else {
		fmt.Printf("Symbol %v is not found ", symbol)
		return false, nil
	}

}

// CreateUrl generates the API endpoint URL
func (s *Stocks) createUrl() (string, error) {
	basePath, err := url.JoinPath(baseUrl, "v2/aggs/ticker", s.symbol, "prev")
	if err != nil {
		return "", fmt.Errorf("error joining path: %w", err)
	}

	parsedURL, err := url.Parse(basePath)
	if err != nil {
		return "", fmt.Errorf("error parsing URL: %w", err)
	}

	queryParams := url.Values{}
	queryParams.Set("adjusted", "true")

	parsedURL.RawQuery = queryParams.Encode()
	return parsedURL.String(), nil
}

// findStock fetches and parses exchange rate data
func (s *Stocks) findStock() error {

	fullUrl, err := s.createUrl()
	if err != nil {
		return err
	}

	requestBody, err := s.apiClient.GetRequest(fullUrl)
	if err != nil {
		return fmt.Errorf("API request error: %w", err)
	}

	if err := s.parseStocksResponseBody(requestBody); err != nil {
		return fmt.Errorf("parsing response error: %w", err)
	}

	return nil
}

// GetRate fetches the exchange rate
func (s *Stocks) GetPrice() (float64, error) {
	if err := s.findStock(); err != nil {
		return 0.0, err
	}
	return s.price, nil
}

// parseResponseBody parses JSON response
func (s *Stocks) parseStocksResponseBody(body []byte) error {
	fxParser := JsonParser[StockInformation]{}
	parsedBody, err := fxParser.ParseResponseBody(body)
	if err != nil {
		return fmt.Errorf("error parsing response body: %w", err)
	}

	s.price = parsedBody.Results[0].Open
	return nil
}
