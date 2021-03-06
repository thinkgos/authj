package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/thinkgos/gin-middlewares/gzap"
)

func main() {
	r := gin.New()

	logger, _ := zap.NewProduction()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(gzap.Logger(logger,
		gzap.WithTimeFormat(time.RFC3339),
		gzap.WithUTC(),
		gzap.WithCustomFields(
			func(c *gin.Context) zap.Field { return zap.String("custom field1", c.ClientIP()) },
			func(c *gin.Context) zap.Field { return zap.String("custom field2", c.ClientIP()) },
		),
	))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(gzap.Recovery(logger, true,
		gzap.WithCustomFields(
			func(c *gin.Context) zap.Field { return zap.String("custom field1", c.ClientIP()) },
			func(c *gin.Context) zap.Field { return zap.String("custom field2", c.ClientIP()) },
		),
	))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Example when panic happen.
	r.GET("/panic", func(c *gin.Context) {
		panic("An unexpected error happen!")
	})

	r.GET("/error", func(c *gin.Context) {
		c.Error(errors.New("An error happen 1")) // nolint: errcheck
		c.Error(errors.New("An error happen 2")) // nolint: errcheck
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
