package main

import (
	"os"

	"github.com/spf13/cobra"
)

type issueRequest struct {
	Data Data `json:"data"`
}

type Data struct {
	CryptoCurrencyList []CryptoCurrency `json:"cryptoCurrencyList"`
}

type CryptoCurrency struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	CMCRank int     `json:"cmcRank"`
	Quotes  []Quote `json:"quotes"`
}

type Quote struct {
	Name             string  `json:"name"`
	PercentChange1h  float32 `json:"percentChange1h"`
	PercentChange24h float32 `json:"percentChange24h"`
	PercentChange7d  float32 `json:"percentChange7d"`
}

var rootCmd = &cobra.Command{
	Use:   "gocoin",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	// Add subcommands
	rootCmd.AddCommand(lsCmd)
	// rootCmd.AddCommand(RepositoryCmd)
	// rootCmd.PersistentFlags().StringVarP(&UserName, "username", "u", "", "user name")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
