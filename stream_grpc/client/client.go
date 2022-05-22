package main

import (
	"context"
	"fmt"
	"go_rpc/stream_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	// 																	不做安全检查
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed dial err %v \n", err)
	}
	// 关闭连接
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	r, err := c.GetStream(ctx, &proto.StreamReqData{Data: "客户端"})
	if err != nil {
		log.Fatalf("could not greet: %v \n", err)
	}
	for {
		a, err := r.Recv()
		if err != nil {
			//log.Fatalf("客户端收取消息失败 %v", err)
			fmt.Printf("客户端收取消息失败 %v \n", err)
			break
		}
		fmt.Println(a.Data)
	}
	// 发送客户端流模式
	putS, err := c.PutStream(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v \n", err)
	}
	for i := 0; i < 10; i++ {
		err := putS.Send(&proto.StreamReqData{Data: "客户端发送" + strconv.Itoa(i)})
		if err != nil {
			fmt.Printf("客户端收取消息失败 %v \n", err)

			break
		}
		time.Sleep(time.Second)
	}
	// 双向流
	allS, err := c.AllStream(ctx)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {

			s, err := allS.Recv()
			if err != nil {
				fmt.Printf("client recv err %v \n", err)
				break
			}
			fmt.Println(s.Data)

		}
	}()
	go func() {
		defer wg.Done()
		for {
			err := allS.Send(&proto.StreamReqData{
				Data: "客户端发送服务端数据",
			})
			if err != nil {
				fmt.Printf("client send err %v \n", err)
				break
			}
			time.Sleep(time.Second)

		}
	}()
	wg.Wait()

}
