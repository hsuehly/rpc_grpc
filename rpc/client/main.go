package main

import (
	"fmt"
	"net/rpc"
)

type SayReq struct {
	Name string
}
type SayRep struct {
	Code int
	Data string
}

func main() {

	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("rpc拨号失败err=", err)
	}
	var replay SayRep
	// 远程调用必须知道函数的 call Id
	err = client.Call("helloService.Say", SayReq{Name: "hsuehly"}, &replay)
	if err != nil {
		fmt.Println("调用失败 err", err)
	}
	fmt.Println(replay)
}
