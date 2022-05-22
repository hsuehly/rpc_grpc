package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"

	pb "go_rpc/grpc/proto"
)

//
// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHelloServiceServer
}

// SayHello implements helloworld
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//fmt.Println(ctx.Value("token"))
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
	// 服务端请求拦截
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("拦截到了新的请求")
		// 拦截获取metadata 里的认证信息
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return resp, status.Error(codes.Unauthenticated, "无token信息")
		}
		fmt.Println("md", md)
		var (
			appId  string
			appKey string
		)

		if appIdSlice, ok := md["appid"]; ok {
			//fmt.Println(appId)
			appId = appIdSlice[0]
			//fmt.Println("err 取不到值")
		}
		if appKeySlice, ok := md["appkey"]; ok {
			appKey = appKeySlice[0]
		}
		if appId != "10101" || appKey != "I am Key" {
			return resp, status.Error(codes.Unauthenticated, "t认证信息错误")
		}
		//return handler(ctx, req)
		// 在请求处理完之后，返回给用户之前需要处理的逻辑这样写
		res, err := handler(ctx, req)
		if err != nil {
			return nil, err
		}
		fmt.Println("请求已经完成")
		return res, err

	}
	opt := grpc.UnaryInterceptor(interceptor)
	// 添加中间件，或者拦截器
	s := grpc.NewServer(opt)
	pb.RegisterHelloServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
