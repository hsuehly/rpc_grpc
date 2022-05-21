package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type helloService struct {
}
type SayReq struct {
	Name string
}
type SayRep struct {
	Code int
	Data string
}

func (h *helloService) Say(req SayReq, reply *SayRep) error {
	*reply = SayRep{Code: 200, Data: "hello" + req.Name}
	return nil

}
func main() {
	// 1.实列化一个server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("启动监听失败err=", err)

	}

	// 2.注册逻辑处理 handle
	// 调用时结构体可以不区分大小写，但是方法必须大写
	err = rpc.RegisterName("helloService", &helloService{})
	if err != nil {
		fmt.Println("rep函数注册失败err", err)

	}

	// 3. 启动服务
	for {
		con, err := lis.Accept()
		if err != nil {
			fmt.Println("rpc启动失败 err=", err)
		}
		go process(con)
	}

}

// go 语言序列化和反序列化的协议是（Gob）

func process(con net.Conn) {
	fmt.Println("来了一个连接")
	rpc.ServeConn(con)

}
