package authj

import (
	"context"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// cttAuthKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type cttAuthKey struct{}

// NewAuthorizer returns the authorizer
// uses a Casbin enforcer and Subject function as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		//checks the userName,path,method permission combination from the request.
		allowed, err := e.Enforce(subject(c), c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Permission validation errors occur!",
			})
			c.Abort()
			return
		} else if !allowed {
			// the 403 Forbidden to the client
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Permission denied!",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// subject returns the value associated with this context for subjectCtxKey,
func subject(c *gin.Context) string {
	v, _ := c.Request.Context().Value(cttAuthKey{}).(string)
	return v
}

// ContextWithSubject return a copy of parent in which the value associated with
// subjectCtxKey is subject.
func ContextWithSubject(c *gin.Context, subject string) {
	ctx := context.WithValue(c.Request.Context(), cttAuthKey{}, subject)
	c.Request = c.Request.WithContext(ctx)
}
