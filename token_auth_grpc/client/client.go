package main

import (
	"context"
	"fmt"
	pb "go_rpc/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type customCredential struct {
}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"appid": "10101", "appkey": "I am Key"}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
// 传输要不要安全
func (c customCredential) RequireTransportSecurity() bool {
	return false
}
func main() {
	// 客户端拦截
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		// todo 							坑 放里面的键 全部为为小写，写大写无效 值可以写大小写
		// 传递认证信息 第一种方法 metadata
		//md := metadata.New(map[string]string{"appid": "10101", "appkey": "I am Key"})
		//ctx = metadata.NewOutgoingContext(ctx, md)
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("请求耗时:%s \n", time.Since(start))
		return err
	}
	// 传递认证信息第二种方法
	opt := grpc.WithPerRPCCredentials(customCredential{})
	//opt := grpc.WithUnaryInterceptor(interceptor)
	//conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()), opt)

	// 另一种写法
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(interceptor), opt)
	conn, err := grpc.Dial("localhost:8080", opts...)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// 关闭
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)
	//md := metadata.Pairs("token", "123456")
	response, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "hsuehly",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(response.GetMessage())
}
