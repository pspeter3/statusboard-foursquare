package main

import (
	"net/http"
)

const BaseUrl = "https://api.foursquare.com/v2/"

func main() {
	http.ListenAndServe(":3000", &foursquareHandler{authToken: ""})
}
