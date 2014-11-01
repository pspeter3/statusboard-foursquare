package main

import (
	"net/http"
	"net/url"
)

type httpGetter interface {
	Get(url string) (*http.Response, error)
}

type Client struct {
	baseUrl   string
	authToken string
	version   string
	http      httpGetter
}

func (c *Client) Url(path string, params *url.Values) (string, error) {
	if params == nil {
		params = &url.Values{}
	}
	params.Add("oauth_token", c.authToken)
	params.Add("v", c.version)
	baseUrl, err := url.Parse(c.baseUrl + path)
	baseUrl.RawQuery = params.Encode()
	return baseUrl.String(), err
}

func (c *Client) Recent() {
	url, _ := c.Url("checkins/recent", nil)
	c.http.Get(url)
}
