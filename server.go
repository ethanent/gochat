package main

import (
	"fmt"
	"io"
	"net"

	"github.com/golang/protobuf/proto"

	"github.com/ethanent/gochat/pcol"
	"github.com/ethanent/protostream"
)

type Session struct {
	name   *string
	conn   net.Conn
	stream *protostream.Stream
}

var sessions []Session = []Session{}

func hostServer() {
	ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("> Server listening on :8080")

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	stream := factory.CreateStream()

	stream.Out(conn)

	session := Session{
		name:   nil,
		conn:   conn,
		stream: stream,
	}

	sessions = append(sessions, session)

	stream.Subscribe(&pcol.RegisterClient{}, func(data proto.Message) {
		parsed := data.(*pcol.RegisterClient)

		session.name = &parsed.Name

		for _, sess := range sessions {
			sess.stream.Push(&pcol.RecvBroadcast{
				Sender:  "Server",
				Message: *session.name + " has entered the chat",
			})
		}

		fmt.Println(parsed.Name, "registered")
	})

	stream.Subscribe(&pcol.SendChat{}, func(data proto.Message) {
		parsed := data.(*pcol.SendChat)

		if session.name == nil {
			stream.Push(&pcol.RecvChat{
				Username: "Server",
				Message:  "Failed to send message. Client has not registered.",
			})
			return
		}

		if len(parsed.Message) < 1 {
			stream.Push(&pcol.RecvBroadcast{
				Sender:  "Server",
				Message: "Please type a message.",
			})
			return
		}

		fmt.Println(*session.name, ":", parsed.Message)

		// Broadcast message

		for _, sess := range sessions {
			sess.stream.Push(&pcol.RecvChat{
				Username: *session.name,
				Message:  parsed.Message,
			})
		}
	})

	// Start copying conn data to stream

	_, err := io.Copy(stream, conn)

	// Conn ended probably

	fmt.Println("Terminating conn", err)

	conn.Close()

	// Remove session from list

	for i, sess := range sessions {
		if sess == session {
			sessions = append(sessions[0:i], sessions[i+1:]...)
			fmt.Println("Session dropped successfully!")
			return
		}
	}

	fmt.Println("Session drop failed. Did not locate session.")
}
