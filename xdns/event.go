package xdns

import (
	"github.com/miekg/dns"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Event struct {
	Name            string
	TypeName        string
	ResponseTargets []string
}

func NewEvent(msg *dns.Msg) Event {
	q := msg.Question[0]

	event := Event{
		Name:            q.Name,
		TypeName:        dns.Type(q.Qtype).String(),
		ResponseTargets: make([]string, 0),
	}

	for _, a := range msg.Answer {
		switch ans := a.(type) {
		case *dns.A:
			event.ResponseTargets = append(event.ResponseTargets, ans.A.String())
		}
	}

	return event
}

func (e Event) getResponseTargets() string {
	return strings.Join(e.ResponseTargets, ",")
}

func (e Event) PrintLog() {

	fields := log.Fields{
		"Name":            e.Name,
		"TypeName":        e.TypeName,
		"ResponseTargets": e.getResponseTargets(),
	}

	log.WithFields(fields).Info("[dns]")
}
