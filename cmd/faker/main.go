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
	c := make(chan core.Event)
	go consumEvents(c)

	httpServer := web.NewServer(c)
	httpServer.Start("0.0.0.0", 8080)

}
