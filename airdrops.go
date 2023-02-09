package main

import "encoding/json"

type ListAirdropResponse struct {
	Data struct {
		Projects []Project `json:"projects"`
	} `json:"data"`
}

type Project struct {
	ProjectName string `json:"projectName"`
}

func (c *Client) GetAirdrops(status string) ([]Project, error) {
	res, err := c.Get(c.endpoint + "/airdrop/query?status=" + status + "&start=0&limit=10&platform=web")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	var r ListAirdropResponse
	decoder.Decode(&r)

	return r.Data.Projects, nil
}
