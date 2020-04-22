package pprof

import (
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
		group.GET("/", gin.WrapF(pprof.Index))
		group.GET("/cmdline", gin.WrapF(pprof.Cmdline))
		group.GET("/profile", gin.WrapF(pprof.Profile))
		group.POST("/symbol", gin.WrapF(pprof.Symbol))
		group.GET("/symbol", gin.WrapF(pprof.Symbol))
		group.GET("/trace", gin.WrapF(pprof.Trace))
		group.GET("/allocs", gin.WrapH(pprof.Handler("allocs")))
		group.GET("/block", gin.WrapH(pprof.Handler("block")))
		group.GET("/goroutine", gin.WrapH(pprof.Handler("goroutine")))
		group.GET("/heap", gin.WrapH(pprof.Handler("heap")))
		group.GET("/mutex", gin.WrapH(pprof.Handler("mutex")))
		group.GET("/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
	}
}
