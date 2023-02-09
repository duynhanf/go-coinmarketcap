package main

import (
	"fmt"
	"os"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocoin",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var lsCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewClient("https://api.coinmarketcap.com/data-api/v3")

		coins, err := client.ListCoins()
		if err != nil {
			return
		}
		table := uitable.New()
		table.MaxColWidth = 50

		table.AddRow("#", "NAME", "SYMBOL", "1h %", "24h %", "7d %")
		for _, currency := range coins {
			table.AddRow(
				currency.CMCRank,
				currency.Name,
				currency.Symbol,
				fmt.Sprintf("%.2f", currency.Quotes[0].PercentChange1h),
				fmt.Sprintf("%.2f", currency.Quotes[0].PercentChange24h),
				fmt.Sprintf("%.2f", currency.Quotes[0].PercentChange7d),
			)
		}
		fmt.Println(table)
	},
}

var airdropCmd = &cobra.Command{
	Use:   "airdrop",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := NewClient("https://api.coinmarketcap.com/data-api/v3")

		airdrops, err := client.GetAirdrops("Ended")
		if err != nil {
			return
		}
		table := uitable.New()
		table.MaxColWidth = 50

		table.AddRow("#", "NAME")
		for idx, airdrop := range airdrops {
			table.AddRow(
				idx,
				airdrop.ProjectName,
			)
		}
		fmt.Println(table)
	},
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(airdropCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
