package main

import (
	"fmt"
	"stockexchange/stocks"
)

func main() {

	StocksClient, err := stocks.NewStocks("C:GBPAUD", "fx")
	if err != nil {
		fmt.Println("can not create stocks %w", err)
	}

	fmt.Println(StocksClient.GetPrice())

}
