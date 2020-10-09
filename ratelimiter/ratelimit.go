package ratelimiter

import (
	"github.com/didip/tollbooth/v5"
	"github.com/didip/tollbooth/v5/limiter"
	"github.com/gin-gonic/gin"
)

// RateLimit rate limit
func RateLimit(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Data(httpError.StatusCode, lmt.GetMessageContentType(), []byte(httpError.Message))
			c.Abort()
			return
		}
		c.Next()
	}
}
