package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"os/user"

	"github.com/fatih/color"

	"github.com/ethanent/protostream"

	"github.com/ethanent/gochat/pcol"
	"github.com/golang/protobuf/proto"
)

func connServer(addr string) {
	// If unset connAddr, have user choose one

	if addr == "undefined" {
		fmt.Println("Input connAddr:")
		_, err := fmt.Scanln(&addr)

		if err != nil {
			panic(err)
		}

		fmt.Println("Connecting...")
	}

	// Connect

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println("Connection failed.")
		panic(err)
	}

	fmt.Println("> Conn established")

	// Set up stream

	stream := factory.CreateStream()

	stream.Out(conn)

	stream.Subscribe(&pcol.RecvChat{}, func(msg proto.Message) {
		data := msg.(*pcol.RecvChat)

		color.Set(color.FgCyan)

		fmt.Print(data.Username)

		color.Unset()

		fmt.Println(":", data.Message)
	})

	stream.Subscribe(&pcol.RecvBroadcast{}, func(msg proto.Message) {
		data := msg.(*pcol.RecvBroadcast)

		fmt.Println("RECV BCAST")

		fmt.Print("[")

		color.Set(color.FgBlue)

		fmt.Print(data.Sender)

		color.Unset()

		fmt.Print("] ")

		fmt.Println(data.Message)
	})

	// Register to conn

	registerConn(stream)

	// Begin copying conn data to stream

	go startCopy(stream, conn)

	// Handle user input

	buffer := bufio.NewReader(os.Stdin)

	for {
		line, _, err := buffer.ReadLine()

		if err != nil {
			panic(err)
		}

		stream.Push(&pcol.SendChat{
			Message: string(line),
		})
	}
}

func startCopy(stream *protostream.Stream, conn net.Conn) {
	_, err := io.Copy(stream, conn)

	fmt.Println("Connection closed.", err)

	os.Exit(0)
}

func registerConn(stream *protostream.Stream) {
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	stream.Push(&pcol.RegisterClient{
		Name:    user.Name,
		Version: 1,
	})
}
