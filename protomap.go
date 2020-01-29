package main

import (
	"github.com/ethanent/gochat/pcol"
	"github.com/golang/protobuf/proto"
)

func getProto(id int) proto.Message {
	switch id {
	case 0:
		return &pcol.RegisterClient{}
	}

	panic("Unknown message id")
}
