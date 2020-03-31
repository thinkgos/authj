package pprof

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"

	"github.com/thinkgos/gin-middlewares/wrap"
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
		group.GET("/", wrap.HandlerFunc(pprof.Index))
		group.GET("/cmdline", wrap.HandlerFunc(pprof.Cmdline))
		group.GET("/profile", wrap.HandlerFunc(pprof.Profile))
		group.POST("/symbol", wrap.HandlerFunc(pprof.Symbol))
		group.GET("/symbol", wrap.HandlerFunc(pprof.Symbol))
		group.GET("/trace", wrap.HandlerFunc(pprof.Trace))
		group.GET("/allocs", wrap.Handler(pprof.Handler("allocs")))
		group.GET("/block", wrap.Handler(pprof.Handler("block")))
		group.GET("/goroutine", wrap.Handler(pprof.Handler("goroutine")))
		group.GET("/heap", wrap.Handler(pprof.Handler("heap")))
		group.GET("/mutex", wrap.Handler(pprof.Handler("mutex")))
		group.GET("/threadcreate", wrap.Handler(pprof.Handler("threadcreate")))
	}
}
