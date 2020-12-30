package main

import (
	"fmt"
	"net/http"

	"github.com/liuyong-go/yong/core"
)

func main() {
	r := core.NewRouter()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("请求首页")
	})
	r.GET("/test", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("请求test")
	})
	r.Run(":9999")
}
