package main

import (
	"fmt"
	"net/http"
)

type foursquareHandler struct {
	authToken string
}

func (f *foursquareHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Hello from foursquareHandler")
}
