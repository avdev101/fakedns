package xdns

import (
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
)

type Event struct {
	Name     string
	TypeName string
}

func NewEvent(msg *dns.Msg) Event {
	q := msg.Question[0]

	return Event{
		Name:     q.Name,
		TypeName: dns.Type(q.Qtype).String(),
	}
}

func (e Event) PrintLog() {

	fields := log.Fields{
		"Name":     e.Name,
		"TypeName": e.TypeName,
	}

	log.WithFields(fields).Info("[dns]")
}
