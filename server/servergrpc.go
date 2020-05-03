package server

import (
	"fmt"
	"github.com/mikeletux/grpc_test/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcServer struct {
	config Config
}

func NewgRPCServer(config Config) Server {
	return &grpcServer{config}
}

func (s *grpcServer) Serve() error {
	addr := fmt.Sprintf("%v:%v", s.config.Host, s.config.Port)
	ln, err := net.Listen(s.config.Protocol, addr)
	if err != nil {
		return err
	}
	log.Printf("Listening on %v\n", addr)

	srv := grpc.NewServer()
	proto.RegisterCommunicationServer(srv, &grpcCommServer{})
	if err := srv.Serve(ln); err != nil {
		return err
	}

	return nil
}
