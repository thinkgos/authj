package expvar

import (
	"expvar"

	"github.com/gin-gonic/gin"

	"github.com/thinkgos/gin-middlewares/wrap"
)

// Handler wrap expvar.Handler
func Handler() gin.HandlerFunc {
	return wrap.Handler(expvar.Handler())
}
