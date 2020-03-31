package wrap

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTTP(handler func(next http.Handler) http.Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		})).ServeHTTP(c.Writer, c.Request)
	}
}

func HTTPf(handler func(next http.HandlerFunc) http.HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		handler(func(w http.ResponseWriter, r *http.Request) {
			c.Request = r
			c.Next()
		}).ServeHTTP(c.Writer, c.Request)
	}
}

func HandlerFunc(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func Handler(h http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
