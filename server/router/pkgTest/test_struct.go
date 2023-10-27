package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStructRouter struct {
}

// InitTestStructRouter 初始化 测试结构 路由信息
func (s *TestStructRouter) InitTestStructRouter(Router *gin.RouterGroup) {
	tsRouter := Router.Group("ts").Use(middleware.OperationRecord())
	tsRouterWithoutRecord := Router.Group("ts")
	var tsApi = v1.ApiGroupApp.PkgTestApiGroup.TestStructApi
	{
		tsRouter.POST("createTestStruct", tsApi.CreateTestStruct)             // 新建测试结构
		tsRouter.DELETE("deleteTestStruct", tsApi.DeleteTestStruct)           // 删除测试结构
		tsRouter.DELETE("deleteTestStructByIds", tsApi.DeleteTestStructByIds) // 批量删除测试结构
		tsRouter.PUT("updateTestStruct", tsApi.UpdateTestStruct)              // 更新测试结构
	}
	{
		tsRouterWithoutRecord.GET("findTestStruct", tsApi.FindTestStruct)       // 根据ID获取测试结构
		tsRouterWithoutRecord.GET("getTestStructList", tsApi.GetTestStructList) // 获取测试结构列表
	}
}
