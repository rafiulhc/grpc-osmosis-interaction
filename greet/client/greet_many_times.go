package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)


func DoGreetManyTimes(client pb.GreetServiceClient, name string) {
	println("GreetManyTimes called")
	stream, err := client.GreetManyTimes(context.Background(), &pb.GreetRequest{Name: name})
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.Message)
	}
}