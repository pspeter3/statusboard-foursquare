package main

import (
	"html/template"
	"net/http"
	"time"
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
	for _, checkin := range *checkins {
		if _, ok := f.Ids[checkin.User.Id]; ok {
			checkinRows = append(checkinRows, toCheckinRow(checkin))
		}
	}
	err = templates.ExecuteTemplate(rw, "panel.html", &checkinRows)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func toCheckinRow(checkin Checkin) CheckinRow {
	photo := checkin.User.Photo.Prefix + "64x64" + checkin.User.Photo.Suffix
	time := relative(time.Unix(checkin.CreatedAt, 0))
	return CheckinRow{Photo: photo, Venue: checkin.Venue.Name, Time: time}
}
