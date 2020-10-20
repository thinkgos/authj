# requestid

[![GoDoc](https://godoc.org/github.com/thinkgos/gin-middlewares/requestid?status.svg)](https://godoc.org/github.com/thinkgos/gin-middlewares/requestid)

requestid is an requestId(traceId) middleware for [Gin](https://github.com/gin-gonic/gin)

## format `hostname-pid-initrandvalue-sequence`

## Installation

```bash
    go get github.com/thinkgos/gin-middlewares/requestid
```

## Simple Example

```Go
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
```

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
