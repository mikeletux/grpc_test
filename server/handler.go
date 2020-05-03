package server

import (
	"context"
	"github.com/mikeletux/grpc_test/proto"
	"log"
)

type grpcCommServer struct{}

func (s *grpcCommServer) SendMessage(c context.Context, text *proto.Text) (ok *proto.Ok, err error) {
	log.Printf("Received: %v\n", text.GetText())
	return &proto.Ok{Ok: "Message arrived"}, nil
}
