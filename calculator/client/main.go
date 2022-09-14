package main

import (
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50052"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewSumServiceClient(conn)

	DoSum(client)

}