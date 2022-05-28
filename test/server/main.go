package main

import (
	"context"
	"go_rpc/test/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "你好我是go发送的" + req.GetName(),
	}, nil

}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to lister %v \n", err)

	}
	defer lis.Close()
	s := grpc.NewServer()
	proto.RegisterHelloServiceServer(s, &server{})
	log.Printf("端口%v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("filed to server %v \n", err)
	}
}
