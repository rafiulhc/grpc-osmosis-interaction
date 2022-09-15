package main

import (
	"io"
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	println("LongGreet called")
	result := "Hello "
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we've finished reading the client stream
			return stream.SendAndClose(&pb.GreetResponse{
				Message: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		result += req.Name + " "
	}
}