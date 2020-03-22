package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

const (
	// defaultPrefix url prefix of pprof
	defaultPrefix = "/debug/pprof"
)

// Router the standard HandlerFuncs from the net/http/pprof package with
// the provided gin.Engine. prefixOptions is a optional. If not prefixOptions,
// the default path prefix("/debug/pprof") is used, otherwise first prefixOptions will be path prefix.
func Router(router *gin.Engine, prefixOptions ...string) {
	prefix := defaultPrefix
	if len(prefixOptions) > 0 {
		prefix = prefixOptions[0]
	}

	group := router.Group(prefix)
	{
		group.GET("/", handler(pprof.Index))
		group.GET("/cmdline", handler(pprof.Cmdline))
		group.GET("/profile", handler(pprof.Profile))
		group.POST("/symbol", handler(pprof.Symbol))
		group.GET("/symbol", handler(pprof.Symbol))
		group.GET("/trace", handler(pprof.Trace))
		group.GET("/allocs", handler(pprof.Handler("allocs").ServeHTTP))
		group.GET("/block", handler(pprof.Handler("block").ServeHTTP))
		group.GET("/goroutine", handler(pprof.Handler("goroutine").ServeHTTP))
		group.GET("/heap", handler(pprof.Handler("heap").ServeHTTP))
		group.GET("/mutex", handler(pprof.Handler("mutex").ServeHTTP))
		group.GET("/threadcreate", handler(pprof.Handler("threadcreate").ServeHTTP))
	}
}

func handler(h http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
