package main

import (
	"log"
	"net"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50052"

type Server struct{
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}
	log.Printf("Server is listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}