package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "go_rpc/rpc/proto"
)

//type hello struct {
//	Name    string   `json:"name"`
//	Age     int32    `json:"age"`
//	Courses []string `json:"courses"`
//}

func main() {
	req := pb.HelloRequest{
		Name:    "hsuehly",
		Age:     12,
		Courses: []string{"go", "gin", "grpc"},
	}
	//req2 := hello{
	//	Name:    "hsuehly",
	//	Age:     12,
	//	Courses: []string{"go", "gin", "grpc"},
	//}
	reqData, err := proto.Marshal(&req)
	if err != nil {
		return
	}
	reqData2, err := json.Marshal(req)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("reqData2 的值%v 长度 %v \n", reqData2, len(reqData2))
	fmt.Printf("reqData 的值%v 长度 %v \n", reqData, len(reqData))
	var reqData3 pb.HelloRequest
	err = proto.Unmarshal(reqData, &reqData3)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(reqData3.Name, reqData3.Courses)
}
