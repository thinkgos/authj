## gin-middlewares

middleware for gin

[![GoDoc](https://godoc.org/github.com/thinkgos/gin-middlewares?status.svg)](https://godoc.org/github.com/thinkgos/gin-middlewares)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/gin-middlewares?tab=doc)
[![Build Status](https://travis-ci.org/thinkgos/gin-middlewares.svg)](https://travis-ci.org/thinkgos/gin-middlewares)
[![codecov](https://codecov.io/gh/thinkgos/gin-middlewares/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/gin-middlewares)
![Action Status](https://github.com/thinkgos/gin-middlewares/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/gin-middlewares)](https://goreportcard.com/report/github.com/thinkgos/gin-middlewares)
[![Licence](https://img.shields.io/github/license/thinkgos/gin-middlewares)](https://raw.githubusercontent.com/thinkgos/gin-middlewares/master/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/gin-middlewares)](https://github.com/thinkgos/gin-middlewares/tags)

# middleware

- [authj](#authj) is an authorization middleware, it's based on [casbin](https://github.com/casbin/casbin).
- [expvar](#expvar) is handler wrap expvar.Handler.
- [gzap](#gzap) is gzap provides log handling using zap package.
- [nocache](#nocache) noCache is a simple piece of middleware that sets a number of HTTP headers to prevent a router (or subrouter) from being cached by an upstream proxy and/or client.
- [pprof](#pprof) the standard HandlerFuncs from the net/http/pprof package with the provided gin.Engine.
- [ratelimit](#ratelimit) RateLimit rate limit
- [requestid](#requestid) is a middleware that injects a request ID into the context of each request.
- [traceid](#traceid) traceid is a middleware that injects a trace ID into the context of each request. A trace ID is a string of uuid.
