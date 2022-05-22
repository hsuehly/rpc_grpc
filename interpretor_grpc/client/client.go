package main

import (
	"context"
	"fmt"
	pb "go_rpc/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func main() {
	// 客户端拦截
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("请求耗时:%s \n", time.Since(start))
		return err
	}
	//opt := grpc.WithUnaryInterceptor(interceptor)
	//conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)

	// 另一种写法
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial("localhost:8080", opts...)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 关闭
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	//md := metadata.Pairs("token", "123456")
	md2 := metadata.New(map[string]string{"token": "2345666", "pwd": "123444"})
	// 创建一个新的上下文， 原理应该是用了 context.WithValue() 应该包装了一层

	ctx := metadata.NewOutgoingContext(context.Background(), md2)
	response, err := c.SayHello(ctx, &pb.HelloRequest{
		Name: "hsuehly",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(response.GetMessage())
}
