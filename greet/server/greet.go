package main

import (
	"context"
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %s", req)
	return &pb.GreetResponse{
		Message: "Hello " + req.Name,
	}, nil

}