package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RpcResponse struct {
	Id     int         `json:"id"`
	Result string      `json:"result"`
	Error  interface{} `json:"error"`
}

func main() {
	var postBody = make(map[string]interface{})
	var reqSlice = make([]string, 1)

	reqSlice[0] = "hsuehly"
	fmt.Println("reqSlice", reqSlice)
	postBody["method"] = "HelloService.Say"
	postBody["params"] = reqSlice
	postBody["id"] = 1
	postData, err := json.Marshal(postBody)
	if err != nil {
		fmt.Println("json 格式化失败err", err)

	}
	fmt.Println("postbody", string(postData))
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/jsonrpc", bytes.NewBuffer(postData))
	if err != nil {
		fmt.Println("实列请求失败err=", err)
		return
	}

	//request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	req, err := client.Do(request)
	if err != nil {
		fmt.Println("发送请求失败err= ", err)

	}
	defer req.Body.Close()
	reqData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("读取数据失败", err)
	}
	fmt.Println("结果", string(reqData))

	var response RpcResponse
	err = json.Unmarshal(reqData, &response)
	if err != nil {
		fmt.Println("反序列化失败", err)
	}
	fmt.Println("结果", response.Result)
}
