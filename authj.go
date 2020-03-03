package authj

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// jwtAuthorizer stores the casbin handler
type jwtAuthorizer struct {
	*casbin.Enforcer
	Subject
}

// Subject function get subject
type Subject func(c *gin.Context) string

// NewtAuthorizer returns the authorizer
// uses a Casbin enforcer and Subject function as input
func NewtAuthorizer(e *casbin.Enforcer, s Subject) gin.HandlerFunc {
	jwt := &jwtAuthorizer{e, s}
	return func(c *gin.Context) {
		//checks the userName,path,method permission combination from the request.
		allowed, err := jwt.Enforce(jwt.Subject(c), c.Request.URL.Path, c.Request.Method)
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
