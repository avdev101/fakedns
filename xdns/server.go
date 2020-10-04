package xdns

import (
	"fmt"
	"github.com/eremeevdev/faker/core"
	"github.com/miekg/dns"
	"time"
)

type Server struct {
	events chan core.Event
	ttlmap TTLMap
}

func NewServer() Server {
	return Server{
		events: make(chan core.Event),
		ttlmap: NewTTLMap(5 * time.Second),
	}
}

func (s *Server) Events() chan core.Event {
	return s.events
}

func (s *Server) notify(msg *dns.Msg) {
	event := NewEvent(msg)
	s.events <- event

}

func (s *Server) createA(ip string, name string) dns.RR {
	rr, err := dns.NewRR(fmt.Sprintf("%v A %v", name, ip))

	if err != nil {
		panic(err)
	}

	rr.Header().Ttl = 1

	return rr
}

func (s *Server) getSimpleAAnswer(ips []string, name string) []dns.RR {
	result := make([]dns.RR, len(ips))

	for i, ip := range ips {
		result[i] = s.createA(ip, name)
	}

	return result
}

func (s *Server) getSchemeAAnswer(ips []string, scheme []int, name string) []dns.RR {
	result := make([]dns.RR, 1)

	count := s.ttlmap.Get(name)
	position := scheme[count]
	ip := ips[position]
	result[0] = s.createA(ip, name)

	nextCount := count + 1

	if nextCount == len(scheme) {
		nextCount = 0
	}

	s.ttlmap.Set(name, nextCount)

	return result
}

func (s *Server) getAAnswer(q dns.Question) []dns.RR {

	ips := getIps(q.Name)
	scheme := getScheme(q.Name)

	if len(scheme) != 0 {
		return s.getSchemeAAnswer(ips, scheme, q.Name)
	}

	return s.getSimpleAAnswer(ips, q.Name)
}

func (s *Server) getAnswer(msg *dns.Msg) []dns.RR {
	answer := make([]dns.RR, 0)

	for _, q := range msg.Question {
		if q.Qtype == dns.TypeA {
			for _, ans := range s.getAAnswer(q) {
				answer = append(answer, ans)
			}
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
