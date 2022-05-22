package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go_rpc/grpc/proto"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello implements helloworld
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//fmt.Println(ctx.Value("token"))
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("get metadata err")
	}
	if tokenSlice, ok := md["token"]; !ok {
		fmt.Println("没有这个值")

	} else {
		fmt.Println("token", tokenSlice[0])
	}
	//for k, v := range md {
	//	fmt.Printf("md[%v]:%v \n", k, v)
	//	for k2, v2 := range v {
	//		fmt.Println(k2, v2)
	//
	//	}
	//}
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
