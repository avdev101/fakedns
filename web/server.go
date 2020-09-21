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
	events        chan core.Event
	eventsStream  chan core.Event
	wsConnections []*websocket.Conn
}

// NewServer creates new web server
func NewServer(eventsStream chan core.Event) Server {
	return Server{
		events:        make(chan core.Event),
		eventsStream:  eventsStream,
		wsConnections: make([]*websocket.Conn, 0),
	}
}

// Events returns events channel
func (s *Server) Events() chan core.Event {
	return s.events
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

	s.wsConnections = append(s.wsConnections, conn)
}

func (s *Server) logHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/ws.html")
}

func (s *Server) handleEvents() {
	for event := range s.eventsStream {

		for _, conn := range s.wsConnections {

			err := conn.WriteJSON(&event)

			if err != nil {
				log.Printf("can't write json: %v", err)
			}
		}
	}
}

// Start new server
func (s *Server) Start(host string, port int) {
	http.HandleFunc("/ws", s.wsHandler)
	http.HandleFunc("/log", s.logHandler)
	http.HandleFunc("/", s.defaultHandler)

	addr := fmt.Sprintf("%v:%v", host, port)

	go s.handleEvents()

	http.ListenAndServe(addr, nil)
}
