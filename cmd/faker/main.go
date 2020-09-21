package main

import (
	"github.com/eremeevdev/faker/core"
	"github.com/eremeevdev/faker/web"
)

func consumEvents(c chan core.Event) {
	for event := range c {
		event.PrintLog()
	}
}

func main() {
	eventsStream := make(chan core.Event)

	httpServer := web.NewServer(eventsStream)

	serverEvents := httpServer.Events()

	go func() {
		for event := range serverEvents {
			event.PrintLog()
			eventsStream <- event
		}
	}()

	httpServer.Start("0.0.0.0", 8000)

}
