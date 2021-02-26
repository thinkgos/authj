package wrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTP wrap func(next http.Handler) http.Handler
func HTTP(handler func(http.Handler) http.Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	}
}

// HTTPf wrap func(next http.HandlerFunc) http.HandlerFunc
func HTTPf(handler func(http.HandlerFunc) http.HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		}).ServeHTTP(c.Writer, c.Request)
	}
}
