package main

import (
	"fmt"
	"net"

	"github.com/ethanent/gochat/pcol"
	"github.com/golang/protobuf/proto"
)

type Session struct {
	name *string
	conn net.Conn
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
	session := Session{
		name: nil,
		conn: conn,
	}

	frb := []byte{}

	for {
		rb := make([]byte, 8)
		c, err := conn.Read(rb)

		if err != nil {
			// Drop session

			for idx, sess := range sessions {
				if sess == session {
					sessions = append(sessions[:idx], sessions[idx+1:]...)
					break
				}
			}

			fmt.Println("Dropping conn", err)

			return
		}

		frb = append(frb, rb[:c]...)

		// Parse frame if possible

		f := pcol.Frame{}

		err = proto.Unmarshal(frb, &f)

		if err != nil {
			// Frame not fully buffered
			continue
		}

		// Clear read from main buffer

		frb = frb[:f.XXX_Size()]

		// Get a relevant proto struct

		msg := getProto(int(f.MessageId))

		// Parse data

		err = proto.Unmarshal(f.Data, msg)

		if err != nil {
			fmt.Println("Parse error", err)
			continue
		}

		// Handle msg

		fmt.Println(f.MessageId, msg)
	}
}
