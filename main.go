package main

import (
	"github.com/liuyong-go/yong/config"
)

func main() {
	r := config.InitRoute()
	r.Run(":9999")
}
