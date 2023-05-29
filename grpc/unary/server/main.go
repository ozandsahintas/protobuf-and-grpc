package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	p "grpc/unary/proto"
)

type Server struct {
	p.UnaryServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")

	if err != nil {
		log.Fatal("error listening!")
	}

	log.Printf("Listening...")

	s := grpc.NewServer()
	p.RegisterUnaryServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("error serving!")
	}
}

func (s *Server) Unary(ctx context.Context, in *p.UnaryRequest) (*p.UnaryResponse, error) {
	fmt.Println("a request received")

	return &p.UnaryResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
