package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eremeevdev/faker/core"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	// EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

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
func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	event := NewEvent(r)
	s.events <- event
}

func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade: %v", err)
		return
	}

	for event := range s.events {
		err := conn.WriteJSON(&event)
		if err != nil {
			log.Printf("can't write json: %v", err)
		}
	}
}

func (s *Server) logHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/ws.html")
}

// Start new server
func (s *Server) Start(host string, port int) {
	http.HandleFunc("/ws", s.wsHandler)
	http.HandleFunc("/log", s.logHandler)
	http.HandleFunc("/", s.defaultHandler)

	addr := fmt.Sprintf("%v:%v", host, port)
	http.ListenAndServe(addr, nil)
}
