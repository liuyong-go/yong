package core

//RouterGroup 路由组
type RouterGroup struct {
	prefix     string
	middleware []HandlerFunc
	parent     *RouterGroup
	router     *Router
}

//Group 创建分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	router := group.router
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		router: router,
	}
	router.groups = append(router.groups, newGroup)
	return newGroup
}
func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.router.router.addRouter(method, pattern, handler)
}

// GET defines the method to add GET request
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}
