package main

import (
	"context"
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/calculator/proto"
)

func DoSum(client pb.SumServiceClient) {
	log.Printf("DoSum function was invoked")
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	})
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}