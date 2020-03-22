# expvar

[![GoDoc](https://godoc.org/github.com/thinkgos/gin-middlewares/expvar?status.svg)](https://godoc.org/github.com/thinkgos/gin-middlewares/expvar)

```go
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
```

Request: http://localhost:8080/debug/vars