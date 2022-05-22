package main

import (
	"fmt"
	"go_rpc/stream_grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"time"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

// GetStream 服务端流模式
func (s Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		err := res.Send(&proto.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		if err != nil {
			log.Fatalf("failed send err %v \n", err)
			return err
		}
		if i > 10 {
			break
		}
		time.Sleep(time.Second * 1)
	}
	return nil
}

// PutStream 客户端流模式
func (s Server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Printf("服务端收取消息失败 %v \n", err)
			return err
		} else {
			fmt.Println(a.Data)
		}

	}
	return nil
}

// AllStream 双向流模式
func (s Server) AllStream(allStr proto.Greeter_AllStreamServer) error {

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			// 接受客户端发送的流
			a, err := allStr.Recv()
			if err != nil {
				fmt.Printf("服务端 recv err %v \n", err)
				break
			}
			fmt.Println(a.Data)

		}
	}()
	go func() {
		defer wg.Done()

		for {
			err := allStr.Send(&proto.StreamResData{
				Data: "服务端发给客户端数据",
			})
			if err != nil {
				fmt.Printf("服务端 send err %v \n", err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	log.Printf("server listening at %v \n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}
