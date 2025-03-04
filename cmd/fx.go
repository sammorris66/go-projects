/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"stockexchange/stocks"
)

// fxCmd represents the fx command
var fxCmd = &cobra.Command{
	Use:   "fx",
	Short: "Get current fx prices",
	Long: `The "fx" command allows you to fetch real-time foreign exchange (FX) prices 
	for a given currency pair or symbol. This command queries the latest exchange 
	rates and displays the current price.

	Usage examples:

  	# Get the FX price for a specific currency pair
  	fx --symbol=USD/EUR

  	# Using the short flag
  	fx -s GBP/USD

	This command integrates with the stocks package to retrieve FX data and requires 
	a valid symbol to function correctly.`,
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")

		exchange, _ := stocks.GlobalFactory.Create("fx", symbol)

		price, _ := exchange.GetPrice()

		fmt.Println(price)
	},
}

func init() {
	fxCmd.Flags().StringP("symbol", "s", "", "Name of symbol")

}
