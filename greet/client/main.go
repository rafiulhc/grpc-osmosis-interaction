package main

import (
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	// DoGreet(client)
	// DoGreetManyTimes(client, "Rafiul")
	// DoLongGreet(client)
	DoGreetEveryone(client)

}
