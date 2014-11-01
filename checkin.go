package main

import (
	"time"
)

type Photo struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type User struct {
	Id    string `json:"id"`
	Photo Photo  `json:"photo"`
}

type Venue struct {
	Name string `json:"name"`
}

type Checkin struct {
	CreatedAt time.Time `json:"createdAt"`
}

type Response struct {
	Recent []Checkin `json:"recent"`
}

type APIResponse struct {
	Response Response `json:"response"`
}
