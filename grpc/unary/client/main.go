package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	p "grpc/unary/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("error connect! %v", err)
	}

	c := p.NewUnaryServiceClient(conn)

	Unary(c)

	defer conn.Close()
}

func Unary(c p.UnaryServiceClient) {
	fmt.Println("requesting unary server...")

	res, err := c.Unary(context.Background(), &p.UnaryRequest{FirstNumber: 20, SecondNumber: -12})

	if err != nil {
		fmt.Println("unary client failed")
	}

	fmt.Printf("unary server returned: %d\n", res.Result)
}
