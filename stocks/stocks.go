package stocks

import (
	"fmt"
	"log"
	"net/url"
	"slices"
	"strings"
)

const baseUrl string = "https://api.polygon.io"

type ExchangeBase struct {
	symbol    string
	price     float64
	apiClient APIClient
	ticker    Ticker
}

func (e *ExchangeBase) ValidateSymbol(symbol string) (bool, error) {

	resp, err := e.ticker.GetTickers()
	if err != nil {
		return false, fmt.Errorf("can not get the ticker information: %w", err)
	}

	if slices.Contains(resp, symbol) {
		fmt.Printf("Symbol %v is found ", symbol)
		return true, nil
	} else {
		return false, nil
	}

}

// CreateUrl generates the API endpoint URL
func (e *ExchangeBase) createUrl() (string, error) {
	basePath, err := url.JoinPath(baseUrl, "v2/aggs/ticker", e.symbol, "prev")
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
	fmt.Println(parsedURL.String())
	return parsedURL.String(), nil
}

// findStock fetches and parses exchange rate data
func (e *ExchangeBase) findStock() error {

	fullUrl, err := e.createUrl()
	if err != nil {
		return err
	}

	requestBody, err := e.apiClient.GetRequest(fullUrl)
	if err != nil {
		return fmt.Errorf("API request error: %w", err)
	}

	if err := e.parseStocksResponseBody(requestBody); err != nil {
		return fmt.Errorf("parsing response error: %w", err)
	}

	return nil
}

// GetRate fetches the exchange rate
func (e *ExchangeBase) GetPrice() (float64, error) {
	if err := e.findStock(); err != nil {
		return 0.0, err
	}
	return e.price, nil
}

// parseResponseBody parses JSON response
func (e *ExchangeBase) parseStocksResponseBody(body []byte) error {
	fxParser := JsonParser[StockInformation]{}
	parsedBody, err := fxParser.ParseResponseBody(body)
	if err != nil {
		return fmt.Errorf("error parsing response body: %w", err)
	}

	e.price = parsedBody.Results[0].Open
	return nil
}

type Stocks struct {
	ExchangeBase
}

func NewStocks(symbol string) (BaseExchange, error) {
	ticker, err := NewTicker("stocks")
	if err != nil {
		return nil, fmt.Errorf("failed to create ticker: %w", err)
	}

	apiClient, err := NewAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}

	stocks := &Stocks{
		ExchangeBase: ExchangeBase{
			symbol:    symbol,
			apiClient: *apiClient,
			ticker:    *ticker,
		},
	}

	valid, err := stocks.ValidateSymbol(symbol)
	if err != nil {
		log.Fatalf("symbol validation failed: %v", err)
	}
	if !valid {
		log.Fatalf("invalid symbol: %s", symbol)
	}
	return stocks, nil
}

type Fx struct {
	ExchangeBase
}

func NewFx(symbol string) (BaseExchange, error) {
	ticker, err := NewTicker("fx")
	if err != nil {
		return nil, fmt.Errorf("failed to create ticker: %w", err)
	}

	apiClient, err := NewAPIClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create API client: %w", err)
	}

	fx := &Fx{
		ExchangeBase: ExchangeBase{
			symbol:    symbol,
			apiClient: *apiClient,
			ticker:    *ticker,
		},
	}

	// Format symbol correctly
	fx.formatSymbol(symbol)

	valid, err := fx.ValidateSymbol(fx.symbol)
	if err != nil {
		log.Fatalf("symbol validation failed: %v", err)
	}
	if !valid {
		log.Fatalf("invalid symbol: %s", fx.symbol)
	}
	return fx, nil
}

// Custom symbol formatting for FX
func (f *Fx) formatSymbol(symbol string) {
	f.symbol = "C:" + strings.ReplaceAll(symbol, "/", "")
}
