package main

import (
	"flag"

	"github.com/eremeevdev/faker/core"
	"github.com/eremeevdev/faker/web"
	"github.com/eremeevdev/faker/xdns"
)

func getConfig() core.Config {

	var conf core.Config

	flag.Int64Var(&conf.Port, "port", 8000, "port for web server")
	flag.StringVar(&conf.ForceIP, "forceip", "", "force ip addr to use")
	flag.IntVar(&conf.TTL, "ttl", 5, "ttl for scheme")

	flag.Parse()

	return conf
}

func main() {
	eventsStream := make(chan core.Event)

	config := getConfig()

	httpServer := web.NewServer(eventsStream)
	dnsServer := xdns.NewServer(config)

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

	httpServer.Start("0.0.0.0", int(config.Port))
}
