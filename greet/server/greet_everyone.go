package main

import (
	"io"
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	println("GreetEveryone called")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we've finished reading the client stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		result := "Hello " + req.Name + "!"
		err = stream.Send(&pb.GreetResponse{
			Message: result,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}