/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stockexchange",
	Short: "A CLI tool for retrieving stock and foreign exchange (FX) prices.",
	Long: `StockExchange CLI is a command-line tool for retrieving real-time 
	stock market and foreign exchange (FX) data. It provides various commands 
	to fetch stock prices, currency exchange rates, and other financial data.

	Usage examples:

 	# Get FX rates for a currency pair
  	stockexchange fx --symbol=c:USDEUR

  	# Get stock price for a specific company
  	stockexchange stocks --symbol=AAPL

	This tool integrates with external financial data sources to fetch the latest 
	market prices and exchange rates. Use the appropriate subcommands to get 
	the information you need.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Welcome to StockExchange CLI! Use --help to see available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stockexchange.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.AddCommand(fxCmd)
	rootCmd.AddCommand(stocksCmd)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
