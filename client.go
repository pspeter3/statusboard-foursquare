package main

import (
	"net/http"
	"net/url"
)

type httpGetter interface {
	Get(url string) (resp *http.Response, err error)
}

type client struct {
	baseUrl   string
	authToken string
	version   string
	http      httpGetter
}

func (c *client) recent() {
	baseUrl, _ := url.Parse(c.baseUrl + "/checkins/recent")
	params := url.Values{}
	params.Add("oauth_token", c.authToken)
	params.Add("v", c.version)
	baseUrl.RawQuery = params.Encode()
	c.http.Get(baseUrl.RequestURI())
}
