package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func init() {
	//log.SetFormatter(&log.JSONFormatter{})
}

// Event store http event details
type Event struct {
	Port      string
	Path      string
	Host      string
	RawQuery  string
	Query     url.Values
	Method    string
	UserAgent string
	Body      string
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
	e.UserAgent = r.UserAgent()
	e.Body = getBody(r)

	return e
}

func getBody(r *http.Request) string {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return fmt.Sprintf("can't read body: %v", err)
	}

	return string(body)
}

func getPort(r *http.Request) string {
	port := r.Header.Get("X-Forwarded-Port")

	if port == "" {
		port = r.URL.Port()
	}

	return port
}

func getHost(r *http.Request) string {
	host := r.Header.Get("X-Forwarded-Host")

	if host == "" {
		host = r.Host
	}

	return host
}

// PrintLog prints event info
func (e Event) PrintLog() {

	fields := log.Fields{
		"method":    e.Method,
		"host":      e.Host,
		"port":      e.Port,
		"path":      e.Path,
		"query":     e.RawQuery,
		"UserAgent": e.UserAgent,
		"Body":      e.Body,
	}

	log.WithFields(fields).Info("[http]")
}
