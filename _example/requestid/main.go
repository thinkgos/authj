package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/thinkgos/gin-middlewares/requestid"
)

func main() {
	router := gin.New()
	router.Use(requestid.RequestID)
	router.GET("/", func(c *gin.Context) {
		fmt.Println(requestid.FromRequestID(c))
	})
	router.Run(":8080")
}
