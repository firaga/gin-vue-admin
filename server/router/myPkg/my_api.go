package myPkg

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

func (m *MyApi) InitMyApiRouter(Router *gin.RouterGroup) {
	MyRouter := Router.Group("myapi")
	api := v1.ApiGroupApp.MyPkgApiGroup.MyApi
	{
		MyRouter.POST("create", api.CreateApi)
	}
}
