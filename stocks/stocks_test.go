package stocks

import (
	"testing"
	// cSpell:ignore httpmock
)

type MockTicker struct {
	response []string
	err      error
}

func (m *MockTicker) GetTickers() ([]string, error) {
	return m.response, m.err
}

func TestValidateSymbol(t *testing.T) {
	testcases := []struct {
		symbolName string
		expValue   bool
	}{
		{"AAPL", true},
		{"CMCX", false},
		{"C:GBPUSD", true},
		{"INCORRECT", false},
	}
	for _, tc := range testcases {
		t.Run(tc.symbolName, func(t *testing.T) {

			mockTicker := &MockTicker{
				response: []string{"AAPL", "C:GBPUSD"},
				err:      nil,
			}

			ex := &ExchangeBase{
				ticker: mockTicker,
			}

			check, _ := ex.ValidateSymbol(tc.symbolName)

			if check != tc.expValue {
				t.Errorf("check should equal %t", tc.expValue)
			}

		})
	}
}

func TestCreateUrl(t *testing.T) {
	testcases := []struct {
		symbolName string
		expValue   string
	}{
		{"AAPL", "https://api.polygon.io/v2/aggs/ticker/AAPL/prev?adjusted=true"},
	}
	for _, tc := range testcases {
		t.Run(tc.symbolName, func(t *testing.T) {

			ex := ExchangeBase{
				symbol: tc.symbolName,
			}

			url, _ := ex.createUrl()

			if url != tc.expValue {
				t.Errorf("check should equal %s", tc.expValue)
			}

		})
	}
}

func TestGetPrice(t *testing.T) {
	testcases := []struct {
		symbolName string
		expPrice   float64
		expectErr  bool
	}{
		{"AAPL", 1.0, true},
	}
	for _, tc := range testcases {
		t.Run(tc.symbolName, func(t *testing.T) {

			ex := ExchangeBase{
				symbol: tc.symbolName,
			}

			price, err := ex.GetPrice()

			if err == nil && tc.expectErr {
				t.Errorf("There was no error but there should be no error")
			}

			if (price != tc.expPrice) && !tc.expectErr {
				t.Errorf("price should be %f", tc.expPrice)
			}
		})
	}

}
