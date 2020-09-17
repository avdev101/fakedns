package web

import (
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

// Event store http event details
type Event struct {
	Port     int
	Path     string
	Host     string
	RawQuery string
	Query    url.Values
	Method   string
}

// NewEvent creates new event from http request
func NewEvent(r *http.Request) Event {
	var e Event

	e.Method = r.Method
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

	fields := log.Fields{
		"method": e.Method,
		"host":   e.Host,
		"port":   e.Port,
		"path":   e.Path,
		"query":  e.RawQuery,
	}

	log.WithFields(fields).Info("[http]")
}
