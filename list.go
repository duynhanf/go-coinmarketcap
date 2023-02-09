package main

import (
	"encoding/json"
	"fmt"

	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		client := NewClient("https://api.coinmarketcap.com/data-api/v3/cryptocurrency/listing?start=1&limit=10&sortBy=market_cap&sortType=desc&convert=USD&cryptoType=all&tagType=all&audited=false&aux=ath,atl,high24h,low24h,num_market_pairs,cmc_rank,date_added,max_supply,circulating_supply,total_supply,volume_7d,volume_30d,self_reported_circulating_supply,self_reported_market_cap")

		res, err := client.Get(client.endpoint)
		if err != nil {
			return
		}

		defer res.Body.Close()

		decoder := json.NewDecoder(res.Body)
		var r issueRequest
		decoder.Decode(&r)

		table := uitable.New()
		table.MaxColWidth = 50

		table.AddRow("#", "NAME", "SYMBOL", "1h %", "24h %", "7d %")
		for _, currency := range r.Data.CryptoCurrencyList {
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
