package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/zepeng-jiang/Go-000/Week02/internal/pkg/handler"
	"io/ioutil"
	"net/http"
	"testing"
)

// 运行本测试文件，需要将程序启动

var (
	url    = "http://127.0.0.1:8080/user"
	method = "POST"
)

func TestCreateUserHandlerBadPath(t *testing.T) {
	req := &handler.CreateUserRequest{
		ID:       996,
		Name:     "程序员",
		Password: "password",
	}

	result := SendAndAccept(req, method, url)
	if result.Code == 9999 {
		t.Log("满足预期情况")
	} else {
		t.Error("不满足预期情况")
	}
}

func TestCreateUserHandlerHappyPath(t *testing.T) {
	req := &handler.CreateUserRequest{
		ID:       965,
		Name:     "靓仔",
		Password: "password",
	}
	result := SendAndAccept(req, method, url)
	if result.Code == 200 {
		t.Log("满足预期情况")
	} else {
		t.Error("不满足预期情况")
	}
}

func TestCreateUserHandlerSadPath(t *testing.T) {
	req := &handler.CreateUserRequest{
		ID:       0,
		Name:     "路人甲",
		Password: "password",
	}

	result := SendAndAccept(req, method, url)
	if result.Code == 400 {
		t.Log("满足预期情况")
	} else {
		t.Error("不满足预期情况")
	}
}

// SendAndAccept 发送 HTTP 请求
func SendAndAccept(request *handler.CreateUserRequest, method string, url string) *handler.CreateUserResponse {
	// 构造请求
	req := generateRequest(request, method, url)

	// 发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("send request failed, error: ", err.Error())
		return nil
	}
	defer resp.Body.Close()

	return generateResponse(resp)
}

// generateRequest 构造请求
func generateRequest(request *handler.CreateUserRequest, method string, url string) *http.Request {
	reqData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("marshal request failed! error: ", err.Error())
		return nil
	}
	req, err1 := http.NewRequest(method, url, bytes.NewReader([]byte(reqData)))
	if err1 != nil {
		fmt.Println("create user request failed, error: ", err1.Error())
		return nil
	}
	req.Header.Add("Content-Type", "application/json")
	return req
}

// generateResponse 解析响应并发序列化为对应的结构体
func generateResponse(resp *http.Response) *handler.CreateUserResponse {
	response := &handler.CreateUserResponse{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read response body failed! error: ", err.Error())
		return nil
	}
	err = json.Unmarshal(body, response)
	if err != nil {
		fmt.Println("ummarshal response failed! error: ", err.Error())
		return nil
	}
	return &handler.CreateUserResponse{
		Code:    response.Code,
		Message: response.Message,
	}
}
