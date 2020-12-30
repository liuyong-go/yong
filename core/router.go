package core

import (
	"fmt"
	"net/http"
)

//HandlerFunc 请求方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

//Router 声明路由结构体
type Router struct {
	router map[string]HandlerFunc
}

//NewRouter core.Router 结构体引用
func NewRouter() *Router {
	return &Router{router: make(map[string]HandlerFunc)}
}
func (engine *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

//GET get请求
func (engine *Router) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

//POST POST请求
func (engine *Router) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

//PUT PUT请求
func (engine *Router) PUT(pattern string, handler HandlerFunc) {
	engine.addRouter("PUT", pattern, handler)
}

//ServeHttp 实现接口
func (engine *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 not found : %s\n", req.URL)
	}
}

//Run 开启服务监听
func (engine *Router) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
