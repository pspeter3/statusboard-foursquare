package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

type Config struct {
	AuthToken string
	Ids       map[string]bool
	Port      string
	Version   string
}

var config Config

func init() {
	authToken := os.Getenv("AUTH_TOKEN")
	idsString := os.Getenv("IDS")
	port := os.Getenv("PORT")
	version := os.Getenv("VERSION")
	if authToken == "" {
		log.Fatal("No auth token")
	}
	if idsString == "" {
		log.Fatal("No ids")
	}
	if port == "" {
		log.Fatal("No port")
	}
	if version == "" {
		log.Fatal("No version")
	}
	parts := strings.Split(idsString, ",")
	ids := make(map[string]bool)
	for _, id := range parts {
		ids[id] = true
	}
	port = ":" + port
	config = Config{AuthToken: authToken, Ids: ids, Port: port, Version: version}
}

func main() {
	client := &Client{AuthToken: config.AuthToken, Version: config.Version}
	http.ListenAndServe(config.Port, &Handler{Client: client, Ids: config.Ids})
}
