package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	p "grpc/bi-directional/proto"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error connect! %v", err)
	}

	defer conn.Close()

	c := p.NewBiDirectionalServiceClient(conn)
	Communication(c)
}

func Communication(c p.BiDirectionalServiceClient) {
	fmt.Println("communication")

	stream, err := c.Communication(context.Background())

	if err != nil {
		fmt.Println(err)
	}

	reqs := []p.CommRequest{
		{Name: "FirstData"},
		{Name: "SecondData"},
		{Name: "ThirdData"},
		{Name: "FourthData"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Println("Sending: ", &req)
			stream.Send(&req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatal(err)
				break
			}

			fmt.Println("Received: ", res.Result, res.Id)
		}
		close(waitc)
	}()

	<-waitc
}
