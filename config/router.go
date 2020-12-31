package config

import (
	"net/http"

	"github.com/liuyong-go/yong/core"
)

//InitRoute 初始化路由
func InitRoute() *core.Router {
	r := core.NewRouter()
	r.GET("/", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>hello yong</h1>")
	})
	r.GET("/test/*filepath", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s,param:%s", "liuyong", c.Param("filepath"))
	})
	r.GET("/json/:name/:action", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{
			"username": c.Param("name"),
			"action":   c.Param("action"),
		})
	})
	return r
}