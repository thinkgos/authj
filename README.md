# authj

[![GoDoc](https://godoc.org/github.com/thinkgos/authj?status.svg)](https://godoc.org/github.com/thinkgos/authj)
[![Build Status](https://travis-ci.org/thinkgos/authj.svg)](https://travis-ci.org/thinkgos/authj)
[![codecov](https://codecov.io/gh/thinkgos/authj/branch/master/graph/badge.svg)](https://codecov.io/gh/thinkgos/authj)
![Action Status](https://github.com/thinkgos/authj/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/authj)](https://goreportcard.com/report/github.com/thinkgos/authj)
[![Licence](https://img.shields.io/github/license/thinkgos/authj)](https://raw.githubusercontent.com/thinkgos/authj/master/LICENSE)



authj is an authorization middleware for [Gin](https://github.com/gin-gonic/gin), it's based on
 [casbin](https://github.com/casbin/casbin).

## Installation

```bash
go get github.com/thinkgos/authj
```

## Simple Example

```Go
package main

import (
    "net/http"

    "github.com/casbin/casbin/v2"
    "github.com/thinkgos/authj"
    "github.com/gin-gonic/gin"
)

func main() {
    // load the casbin model and policy from files, database is also supported.
    e ,err := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
    if err!= nil{
        panic(err)    
    }   

    // define your router, and use the Casbin authj middleware.
    // the access that is denied by authj will return HTTP 403 error.
    router := gin.New()
    router.Use(authj.NewAuthorizer(e, authj.Subject(func(c *gin.Context) string {
        // return subject like username
        return "admin"
    })))
}
```

## Documentation

The authorization determines a request based on ``{subject, object, action}``, which means what ``subject`` can perform what ``action`` on what ``object``. In this plugin, the meanings are:

1. ``subject``: the logged-on user name
2. ``object``: the URL path for the web resource like "dataset1/item1"
3. ``action``: HTTP method like GET, POST, PUT, DELETE, or the high-level actions you defined like "read-file", "write-blog"

For how to write authorization policy and other details, please refer to [the Casbin's documentation](https://github.com/casbin/casbin).

## Getting Help

- [Casbin](https://github.com/casbin/casbin)
- [Gin](https://github.com/gin-gonic/gin)
- [Gin-authz](https://github.com/gin-contrib/authz)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
