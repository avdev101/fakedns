package main

import (
	"github.com/eremeevdev/faker/core"
	"github.com/eremeevdev/faker/web"
)

func main() {
	c := make(chan core.Event)

	httpServer := web.NewServer(c)
	httpServer.Start("0.0.0.0", 8080)

}
