package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)

func DoGreetEveryone(client pb.GreetServiceClient) {
	println("GreetEveryone called")

	stream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet RPC: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{Name: "Rafiul"},
		{Name: "Annie"},
		{Name: "Asad"},
		{Name: "Ohi"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			println("sending req: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while reading server stream: %v", err)
				break
			}

			log.Printf("Response from GreetEveryone: %v", res.Message)
		}
		close(waitc)
	}()

	<-waitc
}