package core

import (
	"net/http"
	"strings"
)

//HandlerFunc 请求方法
type HandlerFunc func(*Context)

type pather struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newPather() *pather {
	return &pather{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

//解析url生成parts数组，只允许一个*
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *pather) addRouter(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}
func (r *pather) getRouter(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
func (r *pather) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

//Router 声明路由结构体
type Router struct {
	*RouterGroup
	router *pather
	groups []*RouterGroup
}

//NewRouter core.Router 结构体引用
func NewRouter() *Router {
	router := &Router{router: newPather()}
	router.RouterGroup = &RouterGroup{router: router}
	router.groups = []*RouterGroup{router.RouterGroup}
	return router
}
func (engine *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	engine.router.addRouter(method, pattern, handler)
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
	c := newContext(w, req)
	engine.router.handle(c)
}

//Run 开启服务监听
func (engine *Router) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
