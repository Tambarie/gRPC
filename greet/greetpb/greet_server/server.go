package main

import (
	"log"
	"net"
	"github.com/Tambarie/gRPC/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
