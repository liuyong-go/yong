package main

import (
	"net/http"

	"github.com/liuyong-go/yong/core"
)

func main() {
	r := core.NewRouter()
	r.GET("/", func(c *core.Context) {
		c.HTML(http.StatusOK, "<h1>hello yong</h1>")
	})
	r.GET("/test", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s", "liuyong")
	})
	r.GET("/json", func(c *core.Context) {
		c.JSON(http.StatusOK, core.H{
			"username": c.Query("name"),
			"action":   c.Query("action"),
		})
	})
	r.Run(":9999")
}
