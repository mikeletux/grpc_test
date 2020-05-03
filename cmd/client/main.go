package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mikeletux/grpc_test/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

var (
	hostname string
	port     string
	message  string
)

func Init() {
	flag.StringVar(&hostname, "hostname", "localhost", "Address where the client will connect to")
	flag.StringVar(&port, "port", "3333", "Port where the client will connect to")
	flag.StringVar(&message, "m", "Default message", "Message to send to the server")
	flag.Parse()
}

func main() {
	Init()
	addr := fmt.Sprintf("%v:%v", hostname, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Imposible to connect")
	}

	client := proto.NewCommunicationClient(conn)

	//Here we can start using the client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	ok, err := client.SendMessage(ctx, &proto.Text{Text: message})
	if err != nil {
		log.Fatal("Couldn't send the message")
	}
	log.Printf("The message from server was %v", ok.GetOk())
}
