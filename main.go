package main

import (
	"context"
	"log"
	"net"
	"postman-grpc-test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type testServer struct {
	proto.UnimplementedTestServiceServer
}

func (testServer) SayHello(_ context.Context, req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	return &proto.SayHelloResponse{
		Message: "Hello, " + req.GetPerson().GetName(),
	}, nil
}
func (testServer) SayGoodbye(_ context.Context, req *proto.SayGoodbyeRequest) (*proto.SayGoodbyeResponse, error) {
	return &proto.SayGoodbyeResponse{
		Message: "Goodbye, " + req.GetPerson().GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterTestServiceServer(s, &testServer{})
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
