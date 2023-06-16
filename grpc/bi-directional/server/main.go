package main

import (
	"fmt"
	"google.golang.org/grpc"
	p "grpc/bi-directional/proto"
	"io"
	"log"
	"net"
)

type Server struct {
	p.BiDirectionalServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50055")

	if err != nil {
		log.Fatal("error listening!")
	}

	log.Printf("Listening...")

	s := grpc.NewServer()
	p.RegisterBiDirectionalServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("error serving!")
	}
}

func (s *Server) Communication(stream p.BiDirectionalService_CommunicationServer) error {
	fmt.Println("communication")

	var c int32 = 0
	for {

		c = c + 1
		r, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal(err)
		}

		res := "Hello " + r.Name
		err = stream.Send(&p.CommResponse{
			Id:     c,
			Result: res,
		})

		if err != nil {
			log.Fatal(err)
		}
	}

}
