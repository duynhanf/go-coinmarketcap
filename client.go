package main

import "net/http"

type Client struct {
	endpoint string
	*http.Client
}

func NewClient(endpoint string) *Client {
	return &Client{endpoint, http.DefaultClient}
}
