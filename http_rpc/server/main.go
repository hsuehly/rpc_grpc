package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (h *HelloService) Say(req string, reply *string) error {

	*reply = "hello" + req
	return nil

}
func rpcHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("有请求")
	err := rpc.RegisterName("HelloService", &HelloService{})
	if err != nil {
		fmt.Println("rpc注册名字失败err=", err)
	}
	// 自定义连接
	var conn io.ReadWriteCloser = struct {
		io.Writer
		io.ReadCloser
	}{
		ReadCloser: r.Body,
		Writer:     w,
	}

	err = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	if err != nil {
		fmt.Println("rpc.ServeRequestErr", err)
	}

}
func main() {

	http.HandleFunc("/jsonrpc", rpcHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
