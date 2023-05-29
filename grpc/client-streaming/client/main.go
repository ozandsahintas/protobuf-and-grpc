package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	p "grpc/client-streaming/proto"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error connect! %v", err)
	}

	defer conn.Close()

	c := p.NewClientStreamingServiceClient(conn)
	ServerStatus(c)
}

func ServerStatus(c p.ClientStreamingServiceClient) {
	fmt.Println("server id stream starts")

	reqs := []*p.StatusRequest{
		{ServerId: 0},
		{ServerId: 1},
		{ServerId: 2},
	}

	stream, err := c.Status(context.Background())

	if err != nil {
		fmt.Println("error status requests")
	}

	for _, req := range reqs {
		fmt.Printf("server_id %d requested\n", req.ServerId)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		fmt.Printf("error client-stream response: %v", err)
	}

	fmt.Printf("Status: \n%s\n", res.Result)

}
