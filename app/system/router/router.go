package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

func InitSmartApiRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// 无需认证的路由
	smartApiNoCheckRoleRouter(r)
	// 需要认证的路由
	smartApiCheckRoleRouter(r, authMiddleware)
	return r
}

// 无需认证的路由示例
func smartApiNoCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// 需要认证的路由示例
func smartApiCheckRoleRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	for _, f := range routerCheckRole {
		f(v1, authMiddleware)
	}
}
