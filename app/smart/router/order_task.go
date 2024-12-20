// @Author sunwenbo
// 2024/7/12 20:28
package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"smart-api/app/smart/apis"
	"smart-api/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerOrderTaskAuthRouter)
	//routerNoCheckRole = append(routerNoCheckRole, registerOrderRouter)
}

// 注册工单类别路由
func registerOrderTaskAuthRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.OrderTask{}

	r := v1.Group("/order-task").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		// 查询所有的任务
		r.GET("", api.GetPage)
		// 根据ID 查询
		r.GET("/:id", api.Get)
		// 创建任务
		r.POST("", api.Insert)
		// 更新任务信息
		r.PUT("", api.Update)
		// 删除任务
		r.DELETE("", api.Delete)

		// 查询历史任务
		r.GET("/history", api.GetHistoryTaskPage)
		// 删除历史执行后的任务
		r.DELETE("/history", api.DeleteHistoryTask)

		// 查询历史任务日志
		r.GET("/history/:id/logs", api.GetTaskLogsByID)
	}
}
