package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/http-middlewares/requestid"

	"github.com/thinkgos/gin-middlewares/wrap"
)

// RequestID is a middleware that injects a request ID into the context of each
// request. A request ID is a string of the form "host.example.com/random-0001",
// where "random" is a base62 random string that uniquely identifies this go
// process, and where the last number is an atomically incremented request
// counter.
func RequestID() gin.HandlerFunc {
	return wrap.HTTP(requestid.RequestID)
}

// FromRequestID returns a request ID from the given context if one is present.
// Returns the empty string if a request ID cannot be found.
func FromRequestID(c *gin.Context) string {
	return requestid.FromRequestID(c.Request.Context())
}
