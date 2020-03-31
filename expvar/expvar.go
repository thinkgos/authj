package expvar

import (
	"expvar"

	"github.com/gin-gonic/gin"

	"github.com/thinkgos/gin-middlewares/wrap"
)

func Handler() gin.HandlerFunc {
	return wrap.Handler(expvar.Handler())
}
