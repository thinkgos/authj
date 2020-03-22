# pprof

[![GoDoc](https://godoc.org/github.com/thinkgos/gin-middlewares/pprof?status.svg)](https://godoc.org/github.com/thinkgos/gin-middlewares/pprof)

gin pprof middleware

> Package pprof serves via its HTTP server runtime profiling data in the format expected by the pprof visualization tool.

## Usage

### Start using it

Download and install it:

```bash
    go get github.com/gin-middlewares/pprof
```

Import it in your code:

```go
    import "github.com/thinkgos/gin-middlewares/pprof"
```

### Example

```go
package main

import (
    "github.com/thinkgos/gin-middlewares/pprof"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    pprof.Router(router)
    router.Run(":8080")
}
```

### Use the pprof tool

Then use the pprof tool to look at the heap profile:

```bash
    go tool pprof http://localhost:8080/debug/pprof/heap
```

Or to look at a 30-second CPU profile:

```bash
    go tool pprof http://localhost:8080/debug/pprof/profile
```

Or to look at the goroutine blocking profile, after calling runtime.SetBlockProfileRate in your program:

```bash
    go tool pprof http://localhost:8080/debug/pprof/block
```

Or to collect a 5-second execution trace:

```bash
    wget http://localhost:8080/debug/pprof/trace?seconds=5
```
