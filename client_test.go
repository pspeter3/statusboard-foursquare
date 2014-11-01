package main

import (
	"testing"
)

func TestClientUrl(t *testing.T) {
	const (
		baseUrl   = "https://api.foursquare.com/v2/"
		authToken = "token"
		version   = "1"
		path      = "checkins/recent"
	)
	client := Client{baseUrl: baseUrl, authToken: authToken, version: version}
	url, err := client.Url(path, nil)
	if err != nil {
		t.Error(err)
	}
	expected := baseUrl + path + "?oauth_token=" + authToken + "&v=" + version
	if url != expected {
		t.Errorf("Expected %v but got %v", expected, url)
	}
}
