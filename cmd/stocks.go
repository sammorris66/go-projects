package cmd

import (
	"fmt"
	"stockexchange/stocks"

	"github.com/spf13/cobra"
)

// fxCmd represents the fx command
var stocksCmd = &cobra.Command{
	Use:   "stocks",
	Short: "Get current stock prices",
	Long: `The "stock" command allows you to fetch real-time stocks symbol. This command queries the latest exchange 
	rates and displays the current price.

	Usage examples:

  	# Get the price for a specific stock
  	stocks --symbol=AAPL

  	# Using the short flag
  	stocks -s AAPL

	This command integrates with the stocks package to retrieve stock data and requires 
	a valid symbol to function correctly.`,
	Run: func(cmd *cobra.Command, args []string) {
		symbol, _ := cmd.Flags().GetString("symbol")

		exchange, _ := stocks.GlobalFactory.Create("stocks", symbol)

		price, _ := exchange.GetPrice()

		fmt.Println(price)
	},
}

func init() {
	stocksCmd.Flags().StringP("symbol", "s", "", "Name of symbol")

}
