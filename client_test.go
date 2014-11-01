package main

import (
	"testing"
)

func TestClientUrl(t *testing.T) {
	const (
		authToken = "token"
		version   = "1"
		path      = "checkins/recent"
		expected  = BaseUrl + path + "?oauth_token=" + authToken + "&v=" + version
	)
	client := Client{baseUrl: BaseUrl, authToken: authToken, version: version}
	url, err := client.Url(path, nil)
	if err != nil {
		t.Error(err)
	}
	if url != expected {
		t.Errorf("Expected %v but got %v", expected, url)
	}
}
