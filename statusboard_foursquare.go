package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", &foursquareHandler{authToken: ""})
}
