package main

import (
	"fmt"
	"github.com/ethanent/gochat/pcol"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
)

func connServer(addr string) {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		fmt.Println("Connection failed.")
		panic(err)
	}

	fmt.Println("> Conn established")

	go handleIncoming(conn)

	regMsg, err := proto.Marshal(&pcol.RegisterClient{
		Name: "Ethan",
	})

	if err != nil {
		panic(err)
	}

	msgFrame, err := proto.Marshal(&pcol.Frame{
		MessageId: 0,
		Data:      regMsg,
	})

	if err != nil {
		panic(err)
	}

	conn.Write(msgFrame)
}

func handleIncoming(conn net.Conn) {
	ob := []byte{}

	for {
		rb := make([]byte, 8)

		c, err := conn.Read(rb)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Server has closed.")
				panic("Server closed.")
			}

			panic(err)
		}

		ob = append(ob, rb[:c]...)

		f := pcol.Frame{}

		err = proto.Unmarshal(ob, &f)

		if err != nil {
			// Full frame has not been buffered
			continue
		}

		pf := getProto(int(f.MessageId))

		err = proto.Unmarshal(f.Data, pf)

		if err != nil {
			fmt.Println("Server provided malformatted message")
			panic(err)
		}

		// Clear unread portion of ob

		ob = ob[:f.XXX_Size()]

		// Handle msg

		fmt.Println(f.MessageId, pf)
	}
}
