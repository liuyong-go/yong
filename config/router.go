package config

import (
	"net/http"

	"github.com/liuyong-go/yong/core"
	"github.com/liuyong-go/yong/middleware"
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
	v1 := r.Group("/v1")
	v1.Use(middleware.RequestLog)
	{
		v1.GET("/", func(c *core.Context) {
			c.HTML(http.StatusOK, "<h1>hello yong group</h1>")
		})
		v1.GET("/test", func(c *core.Context) {
			c.HTML(http.StatusOK, "<h1>hello yong group test</h1>")
		})
	}

	return r
}
