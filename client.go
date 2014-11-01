package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Photo struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type User struct {
	Id    string `json:"id"`
	Photo *Photo `json:"photo"`
}

type Venue struct {
	Name string `json:"name"`
}

type Checkin struct {
	CreatedAt uint32 `json:"createdAt"`
	User      *User  `json:"user"`
	Venue     *Venue `json:"venue"`
}

type Response struct {
	Recent *[]Checkin `json:"recent"`
}

type APIResponse struct {
	Response *Response `json:"response"`
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
	err = json.Unmarshal(body, &apiResponse)
	return apiResponse.Response.Recent, err
}
