package middleware

import (
	"log"

	"github.com/liuyong-go/yong/core"
)

//RequestLog 记录请求日志
func RequestLog(c *core.Context) {
	log.Println("请求", c.Method, c.Req.URL, c.Params)
}
