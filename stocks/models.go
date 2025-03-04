package stocks

type StockInformation struct {
	Ticker       string        `json:"ticker"`
	QueryCount   int           `json:"queryCount"`
	ResultsCount int           `json:"resultsCount"`
	Adjusted     bool          `json:"adjusted"`
	Results      []StockResult `json:"results"`
	Status       string        `json:"status"`
	RequestID    string        `json:"request_id"`
	Count        int           `json:"count"`
}

type StockResult struct {
	Ticker                 string  `json:"T"`
	Volume                 float64 `json:"v"`
	VolumeWeightedAvgPrice float64 `json:"vw"`
	Open                   float64 `json:"o"`
	Close                  float64 `json:"c"`
	High                   float64 `json:"h"`
	Low                    float64 `json:"l"`
	Timestamp              int64   `json:"t"`
	Trades                 int     `json:"n"`
}

type TickerInformation struct {
	Results []TickerResult `json:"results"`
}

type TickerResult struct {
	Ticker             string `json:"ticker"`
	Name               string `json:"name"`
	Market             string `json:"market"`
	Locale             string `json:"locale"`
	Active             bool   `json:"active"`
	CurrencySymbol     string `json:"currency_symbol"`
	CurrencyName       string `json:"currency_name"`
	BaseCurrencySymbol string `json:"base_currency_symbol"`
	BaseCurrencyName   string `json:"base_currency_name"`
	LastUpdatedUtc     string `json:"last_updated_utc"`
}
