package expvar

import (
	"expvar"

	"github.com/gin-gonic/gin"
)

// Handler wrap expvar.Handler
func Handler() gin.HandlerFunc {
	return gin.WrapH(expvar.Handler())
}
