package main

import (
	"context"
	"log"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %s", req)
	return &pb.SumResponse{
		Result: req.FirstNumber + req.SecondNumber,
	}, nil

}