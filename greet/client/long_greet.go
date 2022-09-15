package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)

func DoLongGreet(client pb.GreetServiceClient) {
	println("LongGreet called")
	reqs := []*pb.GreetRequest{
		{Name: "Rafiul"},
		{Name: "Hasan"},
		{Name: "Chowdhury"},
	}
	stream, err := client.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}
	for _, req := range reqs {
		println("sending req: ", req.Name)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from LongGreet: %v", err)
	}
	log.Printf("Response from LongGreet: %v", res.Message)
}