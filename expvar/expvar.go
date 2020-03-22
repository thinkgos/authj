package expvar

import (
	"expvar"

	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		expvar.Handler().ServeHTTP(c.Writer, c.Request)
	}
}
