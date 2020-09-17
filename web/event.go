package web

import (
	"log"
	"net/http"
	"net/url"
)

// Event store http event details
type Event struct {
	Port     int
	Path     string
	Host     string
	RawQuery string
	Query    url.Values
}

// NewEvent creates new event from http request
func NewEvent(r *http.Request) Event {
	var e Event

	e.Port = getPort(r)
	e.Path = r.URL.Path
	e.Host = getHost(r)
	e.RawQuery = r.URL.RawQuery
	e.Query = r.URL.Query()

	return e
}

func getPort(r *http.Request) int {
	return 0
}

func getHost(r *http.Request) string {
	return ""
}

// PrintLog prints event info
func (e Event) PrintLog() {
	log.Printf("request")
}
