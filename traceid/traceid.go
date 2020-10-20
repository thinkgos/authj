package traceid

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/http-middlewares/traceid"

	"github.com/thinkgos/gin-middlewares/wrap"
)

// TraceID is a middleware that injects a trace ID into the context of each
// request. A trace ID is a string of uuid.
func TraceID() gin.HandlerFunc {
	return wrap.HTTP(traceid.TraceID)
}

// FromTraceID returns a trace ID from the given context if one is present.
// Returns the empty string if a trace ID cannot be found.
func FromTraceID(c *gin.Context) string {
	return traceid.FromTraceID(c.Request.Context())
}
