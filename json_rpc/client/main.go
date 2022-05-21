package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type SayReq struct {
	Name string
}
type SayRep struct {
	Code int
	Data string
}

func main() {

	//client, err := rpc.Dial("tcp", "localhost:8080")
	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("rpc拨号失败err=", err)
	}
	var replay SayRep
	// 远程调用必须知道函数的 call Id
	//err = client.Call("helloService.Say", SayReq{Name: "hsuehly"}, &replay)
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(con))
	err = client.Call("HelloService.Say", SayReq{Name: "hsuehly"}, &replay)
	if err != nil {
		fmt.Println("调用失败 err", err)
	}
	fmt.Println(replay)
	//转为json 后其他语言也可以调用
	// 格式 {"method":"HelloService.Say","params":"[参数]","id":0}
}
