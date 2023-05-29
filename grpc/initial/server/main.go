package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	p "grpc/initial/proto"
)

type Server struct {
	p.InitialServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")

	if err != nil {
		log.Fatal("error listening!")
	}

	log.Printf("Listening...")

	if err = grpc.NewServer().Serve(lis); err != nil {
		log.Fatal("error serving!")
	}

}
