package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("panel.html"))

type FoursquareClient interface {
	Recent() (*[]Checkin, error)
}

type CheckinRow struct {
	Photo string
	Venue string
	Time  string
}

type Handler struct {
	Client FoursquareClient
	Ids    map[string]bool
}

func (f *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	checkins, err := f.Client.Recent()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	checkinRows := []CheckinRow{}
	var photo string
	for _, checkin := range *checkins {
		if _, ok := f.Ids[checkin.User.Id]; ok {
			photo = checkin.User.Photo.Prefix + "64x64" + checkin.User.Photo.Suffix
			checkinRows = append(checkinRows, CheckinRow{Photo: photo, Venue: checkin.Venue.Name, Time: ""})
		}
	}
	err = templates.ExecuteTemplate(rw, "panel.html", &checkinRows)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}
