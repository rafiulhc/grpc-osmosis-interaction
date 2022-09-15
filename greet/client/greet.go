package main

import (
	"context"
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)

func DoGreet(client pb.GreetServiceClient) {
	log.Printf("doGreet function was invoked with %v", client)
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		Name: "Rafiul",
	})
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Message)
}