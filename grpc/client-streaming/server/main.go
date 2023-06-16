package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"

	p "grpc/client-streaming/proto"
)

type Server struct {
	p.ClientStreamingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")

	if err != nil {
		log.Fatal("error listening!")
	}

	log.Printf("Listening...")

	s := grpc.NewServer()
	p.RegisterClientStreamingServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("error serving!")
	}
}

func (s *Server) Status(stream p.ClientStreamingService_StatusServer) error {
	fmt.Println("sending all status")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&p.StatusResponse{Result: res})
		}

		if err != nil {
			fmt.Println("error reading client-stream")
		}

		res += fmt.Sprintf("%d is %s\n", req.ServerId, strconv.FormatBool(GetStatus(req.ServerId)))
	}
}

func GetStatus(s uint32) bool {
	switch s {
	case 0:
		return true
	case 1:
		return false
	case 2:
		return true
	}
	return false
}
