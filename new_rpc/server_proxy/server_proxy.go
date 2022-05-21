package server_proxy

import (
	"go_rpc/new_rpc/hanlder"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, replay *string) error
}

// 关心的是函数， 鸭子类型， 可以使用接口

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(hanlder.HelloServiceName, srv)
}
