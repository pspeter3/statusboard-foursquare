package main

import (
	"encoding/json"
	"io/ioutil"
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

func (c *Client) Recent() (*[]Checkin, error) {
	url, err := c.Url("checkins/recent", nil)
	if err != nil {
		return nil, err
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var apiResponse APIResponse
	err = json.Unmarshal(body, apiResponse)
	return &apiResponse.Response.Recent, err
}
