package core

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strings"
)

//Recovery 捕获异常
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Print("%s\n\n", trace(message))
				c.HTML(http.StatusOK, "<h1>wrong</h1>")
			}

		}()
		c.Next()
	}

}
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}
