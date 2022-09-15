package main

import (
	"strconv"

	pb "github.com/rafiulhc/grpc-osmosis-interaction/greet/proto"
)

func (s *Server) GreetManyTimes(req *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error{
	println("GreetManyTimes called")
	for i:=0; i<10; i++{
		Message := "Hello " + req.Name + " number " + strconv.Itoa(i)
		res := &pb.GreetResponse{
			Message: Message,
		}
		stream.Send(res)
	}
	return nil
}