package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//H 输出结构
type H map[string]interface{}

//Context 上下文
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	//request info
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
	//middleware
	handlers []HandlerFunc
	index    int
}

//Param 获取参数
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

//Next 执行后续方法
func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

//PostForm 获取post表单提交字段
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

//Query 获取get请求数据
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//Status 设置响应码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

//SetHeader 设置header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

//输出结果，给定模板赋值变量模式
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

//JSON 输出json格式
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode((obj)); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

//Data 输出数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write((data))
}

//HTML 输出html
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
