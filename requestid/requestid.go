package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/http-middlewares/requestid"

	"github.com/thinkgos/gin-middlewares/wrap"
)

// RequestID is a middleware that injects a request ID into the context of each
// request.
func RequestID(opts ...requestid.Option) gin.HandlerFunc {
	return wrap.HTTP(requestid.RequestID(opts...))
}

// FromRequestID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func FromRequestID(c *gin.Context) string {
	return requestid.FromRequestID(c.Request.Context())
}
