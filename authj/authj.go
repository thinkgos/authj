package authj

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/thinkgos/http-middlewares/authj"

	"github.com/thinkgos/gin-middlewares/wrap"
)

// NewAuthorizer returns the authorizer
// uses a Casbin enforcer and Subject function as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	return wrap.HTTPf(authj.NewAuthorizer(e))
}

// ContextWithSubject return a copy of parent in which the value associated with
// subjectCtxKey is subject.
func ContextWithSubject(c *gin.Context, subject string) {
	ctx := authj.ContextWithSubject(c.Request.Context(), subject)
	c.Request = c.Request.WithContext(ctx)
}
