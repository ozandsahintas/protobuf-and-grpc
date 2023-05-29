package main

import (
	"fmt"
	"google.golang.org/grpc"
	p "grpc/server-streaming/proto"
	"log"
	"net"
	"time"
)

type Server struct {
	p.ServerStreamingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")

	if err != nil {
		log.Fatal("error listening!")
	}

	log.Printf("Listening...")

	s := grpc.NewServer()
	p.RegisterServerStreamingServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("error serving!")
	}
}

func (s *Server) Logs(in *p.LogsRequest, stream p.ServerStreamingService_LogsServer) error {
	fmt.Println("database logs requested")
	logs := []string{
		"D", "A", "T", "A", "B", "A", "S", "E", "_", "L", "O", "G", "S",
	}

	for _, e := range logs {
		stream.Send(&p.LogsResponse{DatabaseResult: e})
		time.Sleep(1 * time.Second)
	}

	return nil
}
