package xdns

import (
	"fmt"
	"github.com/miekg/dns"
)

type Event struct {
}

func NewEvent(msg *dns.Msg) Event {
	return Event{}
}

func (e Event) PrintLog() {
	fmt.Println("event")
}
