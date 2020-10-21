package authj

import (
	"context"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation. This technique
// for defining context keys was copied from Go 1.7's new use of context in net/http.
type ctxAuthKey struct{}

// NewAuthorizer returns the authorizer
// uses a Casbin enforcer and Subject function as input
func NewAuthorizer(e casbin.IEnforcer, subject func(c *gin.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// checks the userName,path,method permission combination from the request.
		allowed, err := e.Enforce(subject(c), c.Request.URL.Path, c.Request.Method)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Permission validation errors occur!",
			})
			return
		} else if !allowed {
			// the 403 Forbidden to the client
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "Permission denied!",
			})
			return
		}
		c.Next()
	}
}

// Subject returns the value associated with this context for subjectCtxKey,
func Subject(c *gin.Context) string {
	val, _ := c.Request.Context().Value(ctxAuthKey{}).(string)
	return val
}

// ContextWithSubject return a copy of parent in which the value associated with
// subjectCtxKey is subject.
func ContextWithSubject(c *gin.Context, subject string) {
	ctx := context.WithValue(c.Request.Context(), ctxAuthKey{}, subject)
	c.Request = c.Request.WithContext(ctx)
}
