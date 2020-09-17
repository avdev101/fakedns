package web

import (
	"fmt"
	"net/http"

	"github.com/eremeevdev/faker/core"
)

// Server is a web server
type Server struct {
	events chan core.Event
}

// NewServer creates new web server
func NewServer(events chan core.Event) Server {
	return Server{events: events}
}

// CreateEvent creates Event from http request
func CreateEvent(r *http.Request) Event {
	var e Event

	return e
}

// DefaultHandler is for handle default http requests
func (s *Server) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	event := NewEvent(r)
	s.events <- event
}

// Start new server
func (s *Server) Start(host, port int) {
	http.HandleFunc("/", s.DefaultHandler)

	addr := fmt.Sprintf("%v:%v", host, port)
	http.ListenAndServe(addr, nil)
}
