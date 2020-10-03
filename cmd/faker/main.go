package main

import (
	"github.com/eremeevdev/faker/core"
	"github.com/eremeevdev/faker/web"
	"github.com/eremeevdev/faker/xdns"
)

func main() {
	eventsStream := make(chan core.Event)

	httpServer := web.NewServer(eventsStream)
	dnsServer := xdns.NewServer()

	go func() {
		for event := range dnsServer.Events() {
			event.PrintLog()
			eventsStream <- event
		}
	}()

	go func() {
		for event := range httpServer.Events() {
			event.PrintLog()
			eventsStream <- event
		}
	}()

	go dnsServer.Start()

	httpServer.Start("0.0.0.0", 8000)
}
