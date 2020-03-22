package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/gin-middlewares/pprof"
)

func main() {
	router := gin.Default()
	pprof.Router(router)
	router.Run(":8080")
}
