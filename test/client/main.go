package main

import (
	"context"
	"fmt"
	"go_rpc/test/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("创建连接失败%v \n", err)
	}
	defer conn.Close()
	// 创建grpc 客户端

	c := proto.NewHelloServiceClient(conn)
	res, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: "我是go发送的请求",
	})
	if err != nil {
		log.Fatalf("发送请求失败%v \n", err)
	}
	fmt.Println(res)

}
