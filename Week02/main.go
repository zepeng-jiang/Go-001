package main

import (
	"fmt"
	"github.com/zepeng-jiang/Go-000/Week02/pkg/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/user", handler.CreateUserHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
