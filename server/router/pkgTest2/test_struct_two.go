package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TestStructTwoRouter struct {
}

// InitTestStructTwoRouter 初始化 测试结构二 路由信息
func (s *TestStructTwoRouter) InitTestStructTwoRouter(Router *gin.RouterGroup) {
	tstRouter := Router.Group("tst").Use(middleware.OperationRecord())
	tstRouterWithoutRecord := Router.Group("tst")
	var tstApi = v1.ApiGroupApp.PkgTest2ApiGroup.TestStructTwoApi
	{
		tstRouter.POST("createTestStructTwo", tstApi.CreateTestStructTwo)             // 新建测试结构二
		tstRouter.DELETE("deleteTestStructTwo", tstApi.DeleteTestStructTwo)           // 删除测试结构二
		tstRouter.DELETE("deleteTestStructTwoByIds", tstApi.DeleteTestStructTwoByIds) // 批量删除测试结构二
		tstRouter.PUT("updateTestStructTwo", tstApi.UpdateTestStructTwo)              // 更新测试结构二
	}
	{
		tstRouterWithoutRecord.GET("findTestStructTwo", tstApi.FindTestStructTwo)       // 根据ID获取测试结构二
		tstRouterWithoutRecord.GET("getTestStructTwoList", tstApi.GetTestStructTwoList) // 获取测试结构二列表
	}
}
