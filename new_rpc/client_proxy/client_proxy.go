package client_proxy

import (
	"go_rpc/new_rpc/hanlder"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic(err)
	}
	return HelloServiceStub{conn}

}

func (h HelloServiceStub) Hello(request string, replay *string) error {
	err := h.Call(hanlder.HelloServiceName+".Hello", request, &replay)
	if err != nil {
		return err
	}
	return nil
}
