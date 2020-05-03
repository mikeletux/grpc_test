package main

import (
	"github.com/mikeletux/grpc_test/server"
	"log"
)

func main() {
	srvConfig := server.Config{Protocol: "tcp", Host: "localhost", Port: "3333"}
	srv := server.NewgRPCServer(srvConfig)
	log.Printf("gRPC Server running at %v://%v:%v\n", srvConfig.Protocol, srvConfig.Host, srvConfig.Port)
	log.Fatal(srv.Serve())
}
