// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package authj

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func testAuthjRequest(t *testing.T, router http.Handler, user, path, method string, code int) {
	r, _ := http.NewRequestWithContext(context.TODO(), method, path, nil)
	r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Code != code {
		t.Errorf("%s, %s, %s: %d, supposed to be %d", user, path, method, w.Code, code)
	}
}

func TestBasic(t *testing.T) {
	router := gin.New()
	e, _ := casbin.NewEnforcer("authj_model.conf", "authj_policy.csv")

	router.Use(func(context *gin.Context) {
		ContextWithSubject(context, "alice")
	})
	router.Use(NewAuthorizer(e))
	router.Any("/*anypath", func(c *gin.Context) {
		c.Status(200)
	})

	testAuthjRequest(t, router, "alice", "/dataset1/resource1", "GET", 200)
	testAuthjRequest(t, router, "alice", "/dataset1/resource1", "POST", 200)
	testAuthjRequest(t, router, "alice", "/dataset1/resource2", "GET", 200)
	testAuthjRequest(t, router, "alice", "/dataset1/resource2", "POST", 403)
}

func TestPathWildcard(t *testing.T) {
	router := gin.New()
	e, _ := casbin.NewEnforcer("authj_model.conf", "authj_policy.csv")

	router.Use(func(context *gin.Context) {
		ContextWithSubject(context, "bob")
	})
	router.Use(NewAuthorizer(e))

	router.Any("/*anypath", func(c *gin.Context) {
		c.Status(200)
	})

	testAuthjRequest(t, router, "bob", "/dataset2/resource1", "GET", 200)
	testAuthjRequest(t, router, "bob", "/dataset2/resource1", "POST", 200)
	testAuthjRequest(t, router, "bob", "/dataset2/resource1", "DELETE", 200)
	testAuthjRequest(t, router, "bob", "/dataset2/resource2", "GET", 200)
	testAuthjRequest(t, router, "bob", "/dataset2/resource2", "POST", 403)
	testAuthjRequest(t, router, "bob", "/dataset2/resource2", "DELETE", 403)

	testAuthjRequest(t, router, "bob", "/dataset2/folder1/item1", "GET", 403)
	testAuthjRequest(t, router, "bob", "/dataset2/folder1/item1", "POST", 200)
	testAuthjRequest(t, router, "bob", "/dataset2/folder1/item1", "DELETE", 403)
	testAuthjRequest(t, router, "bob", "/dataset2/folder1/item2", "GET", 403)
	testAuthjRequest(t, router, "bob", "/dataset2/folder1/item2", "POST", 200)
	testAuthjRequest(t, router, "bob", "/dataset2/folder1/item2", "DELETE", 403)
}

func TestRBAC(t *testing.T) {
	router := gin.New()
	e, _ := casbin.NewEnforcer("authj_model.conf", "authj_policy.csv")

	router.Use(func(context *gin.Context) {
		ContextWithSubject(context, "cathy")
	})
	router.Use(NewAuthorizer(e))
	router.Any("/*anypath", func(c *gin.Context) {
		c.Status(200)
	})

	// cathy can access all /dataset1/* resources via all methods because it has the dataset1_admin role.
	testAuthjRequest(t, router, "cathy", "/dataset1/item", "GET", 200)
	testAuthjRequest(t, router, "cathy", "/dataset1/item", "POST", 200)
	testAuthjRequest(t, router, "cathy", "/dataset1/item", "DELETE", 200)
	testAuthjRequest(t, router, "cathy", "/dataset2/item", "GET", 403)
	testAuthjRequest(t, router, "cathy", "/dataset2/item", "POST", 403)
	testAuthjRequest(t, router, "cathy", "/dataset2/item", "DELETE", 403)

	// delete all roles on user cathy, so cathy cannot access any resources now.
	_, _ = e.DeleteRolesForUser("cathy")

	testAuthjRequest(t, router, "cathy", "/dataset1/item", "GET", 403)
	testAuthjRequest(t, router, "cathy", "/dataset1/item", "POST", 403)
	testAuthjRequest(t, router, "cathy", "/dataset1/item", "DELETE", 403)
	testAuthjRequest(t, router, "cathy", "/dataset2/item", "GET", 403)
	testAuthjRequest(t, router, "cathy", "/dataset2/item", "POST", 403)
	testAuthjRequest(t, router, "cathy", "/dataset2/item", "DELETE", 403)
}
