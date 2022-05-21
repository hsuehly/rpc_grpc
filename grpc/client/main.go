package main

import (
	"context"
	"fmt"
	pb "go_rpc/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
	}
	// 关闭
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	response, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "hsuehly",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(response.GetMessage())
}
