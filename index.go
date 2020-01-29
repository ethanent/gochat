package main

import "flag"

var hostMode *bool = flag.Bool("hostserver", false, "Host a server?")
var connAddr *string = flag.String("host", "0.0.0.0:8080", "Host to connect to")

func main() {
	flag.Parse()

	if *hostMode {
		// Host a server

		hostServer()
	} else {
		connServer(*connAddr)
	}
}
