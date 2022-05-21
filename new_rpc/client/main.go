package main

import (
	"fmt"
	"go_rpc/new_rpc/client_proxy"
)

func main() {

	//client, err := rpc.Dial("tcp", "localhost:8080")
	//if err != nil {
	//	fmt.Println("rpc拨号失败err=", err)
	//}
	var replay string
	// 远程调用必须知道函数的 call Id
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:8080")
	//err = client.Call(hanlder.HelloServiceName+".Hello", "hsuehly", &replay)
	err := client.Hello("hsuehly", &replay)
	if err != nil {
		fmt.Println("调用失败 err", err)
	}
	fmt.Println(replay)
}
