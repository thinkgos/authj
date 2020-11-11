package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/gin-middlewares/expvar"
)

func main() {
	router := gin.Default()
	router.GET("/debug/vars", expvar.Handler())
	router.Run(":8080")
}
