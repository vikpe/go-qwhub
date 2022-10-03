package qwhub

import (
	"github.com/go-resty/resty/v2"
	"github.com/vikpe/qw-hub-api/types"
	"github.com/vikpe/serverstat/qserver/mvdsv"
)

type Client struct {
	RestyClient *resty.Client
}

func NewClient() *Client {
	restyClient := resty.New()
	restyClient.SetBaseURL("https://hubapi.quakeworld.nu/v2")

	return &Client{
		RestyClient: restyClient,
	}
}

func (c *Client) GetMvdsvServers(queryParams ...map[string]string) []mvdsv.Mvdsv {
	req := c.RestyClient.R().SetResult([]mvdsv.Mvdsv{})

	if len(queryParams) > 0 {
		req.SetQueryParams(queryParams[0])
	}

	resp, err := req.Get("servers/mvdsv")

	if err != nil || resp.IsError() {
		return make([]mvdsv.Mvdsv, 0)
	}

	servers := resp.Result().(*[]mvdsv.Mvdsv)
	return *servers
}

func (c *Client) GetStreams() []types.TwitchStream {
	req := c.RestyClient.R().SetResult([]types.TwitchStream{})
	res, err := req.Get("streams")

	if err != nil || res.IsError() {
		return make([]types.TwitchStream, 0)
	}

	streams := res.Result().(*[]types.TwitchStream)
	return *streams
}
