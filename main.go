/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"stockexchange/cmd"
	"stockexchange/stocks"
)

func main() {

	stocks.GlobalFactory.Register("stocks", stocks.NewStocks)
	stocks.GlobalFactory.Register("fx", stocks.NewFx)
	cmd.Execute()
}
