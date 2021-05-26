package pricempire

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"net/url"
)

const (
	getAllItems = "https://api.pricempire.com/v1/getAllItems?"
)

type Client struct {
	ApiKey string
}

type Items struct {
	Status bool   `json:"status"`
	Items  []Item `json:"items"`
}

type Item struct {
	Name   string `json:"name"`
	AppId  int    `json:"appId"`
	Prices struct {
		Csgoempire struct {
			Sourceprice float64 `json:"sourcePrice"`
			Price       float64 `json:"price"`
		} `json:"csgoempire"`
		CsgoempireAvg7 struct {
			Sourceprice float64 `json:"sourcePrice"`
			Price       float64 `json:"price"`
		} `json:"csgoempire_avg7"`
		CsgoempireAvg30 struct {
			Sourceprice float64 `json:"sourcePrice"`
			Price       float64 `json:"price"`
		} `json:"csgoempire_avg30"`
		CsgoempireAvg60 struct {
			Sourceprice float64 `json:"sourcePrice"`
			Price       float64 `json:"price"`
		} `json:"csgoempire_avg60"`
	} `json:"prices"`
}

func NewClient(api string) *Client {
	return &Client{
		ApiKey: api,
	}
}

func (c *Client) GetAllItemsBySites(sites string) (*Items, error) {
	request := fasthttp.AcquireRequest()

	requestBody := url.Values{
		"token":  {c.ApiKey},
		"source": {sites},
	}

	request.SetRequestURI(getAllItems + requestBody.Encode())
	request.Header.SetMethod("GET")

	response := fasthttp.AcquireResponse()

	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}

	var items Items
	if err := json.Unmarshal(response.Body(), &items); err != nil {
		return nil, err
	}

	return &items, nil
}
