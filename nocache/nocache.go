package nocache

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/http-middlewares/nocache"

	"github.com/thinkgos/gin-middlewares/wrap"
)

// NoCache is a simple piece of middleware that sets a number of HTTP headers to prevent
// a router (or subrouter) from being cached by an upstream proxy and/or client.
//
// As per http://wiki.nginx.org/HttpProxyModule - NoCache sets:
//      Expires: Thu, 01 Jan 1970 00:00:00 UTC
//      Cache-Control: no-cache, private, max-age=0
//      X-Accel-Expires: 0
//      Pragma: no-cache (for HTTP/1.0 proxies/clients)
func NoCache() gin.HandlerFunc {
	return wrap.HTTP(nocache.NoCache)
}
