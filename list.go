package main

import (
	"encoding/json"
)

type ListCryptoResponse struct {
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

func (c *Client) ListCoins() ([]CryptoCurrency, error) {

	res, err := c.Get(c.endpoint + "/cryptocurrency/listing?start=1&limit=10&sortBy=market_cap&sortType=desc&convert=USD&cryptoType=all&tagType=all&audited=false&aux=ath,atl,high24h,low24h,num_market_pairs,cmc_rank,date_added,max_supply,circulating_supply,total_supply,volume_7d,volume_30d,self_reported_circulating_supply,self_reported_market_cap")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r ListCryptoResponse
	decoder.Decode(&r)
	return r.Data.CryptoCurrencyList, nil
}
