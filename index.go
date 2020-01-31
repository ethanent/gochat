package main

import (
	"flag"
	"fmt"

	"github.com/ethanent/gochat/pcol"
	"github.com/ethanent/protostream"
)

var hostMode *bool = flag.Bool("hostserver", false, "Host a server?")
var connAddr *string = flag.String("host", "undefined", "Host to connect to")

var factory *protostream.Factory

func main() {
	flag.Parse()

	fmt.Println("Gochat")
	fmt.Println("(c) 2020 Ethan Davis")

	initFactory()

	if *hostMode {
		// Host a server

		hostServer()
	} else {
		connServer(*connAddr)
	}
}

func initFactory() {
	factory = protostream.NewFactory()

	factory.RegisterMessage(0, &pcol.RegisterClient{})
	factory.RegisterMessage(1, &pcol.SendChat{})
	factory.RegisterMessage(2, &pcol.RecvChat{})
}
