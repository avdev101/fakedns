package xdns

import (
	"fmt"
	"github.com/eremeevdev/faker/core"
	"github.com/miekg/dns"
	"log"
)

type Server struct {
	events chan core.Event
}

func NewServer() Server {
	return Server{
		events: make(chan core.Event),
	}
}

func (s *Server) Events() chan core.Event {
	return s.events
}

func (s *Server) notify(msg *dns.Msg) {
	event := NewEvent(msg)
	s.events <- event

}

func (s *Server) getAAnswer(q dns.Question) dns.RR {
	addr := "192.168.0.1-192.168.1.1-s123.asdf.ru"
	rr, err := dns.NewRR(fmt.Sprintf("%v A %v", addr, "123.123.123.123"))

	rr.Header().Ttl = 1

	if err != nil {
		log.Fatal("could not parse APAIR record: ", err)
	}

	return rr
}

func (s *Server) getAnswer(msg *dns.Msg) []dns.RR {
	answer := make([]dns.RR, 0)

	for _, q := range msg.Question {
		if q.Qtype == dns.TypeA {
			answer = append(answer, s.getAAnswer(q))
		}
	}

	return answer
}

func (s *Server) defaultHandler(w dns.ResponseWriter, r *dns.Msg) {

	s.notify(r)

	m := new(dns.Msg)
	m.SetReply(r)

	for _, answer := range s.getAnswer(r) {
		m.Answer = append(m.Answer, answer)
	}

	w.WriteMsg(m)
}

func (s *Server) Start() {
	server := &dns.Server{Addr: ":53", Net: "udp"}

	dns.HandleFunc(".", s.defaultHandler)

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
