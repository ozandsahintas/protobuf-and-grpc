package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	p "grpc/server-streaming/proto"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error connect! %v", err)
	}

	c := p.NewServerStreamingServiceClient(conn)

	DatabaseLogs(c)

	defer conn.Close()
}

func DatabaseLogs(c p.ServerStreamingServiceClient) {
	fmt.Println("requesting database logs...")

	stream, err := c.Logs(context.Background(), &p.LogsRequest{ServerId: 22})

	if err != nil {
		log.Fatalf("error stream call: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error stream: %v", err)
		}

		fmt.Print(res.DatabaseResult)
	}

}
