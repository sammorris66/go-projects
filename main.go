package main

import (
	"fmt"
	"stocks/stocks"
)

func main() {

	StocksClient, err := stocks.NewStocks("C:GBPAUD", "fx")
	if err != nil {
		fmt.Println("can not create stocks %w", err)
	}

	StocksClient.ValidateSymbol()
	fmt.Println(StocksClient.GetPrice())

}
