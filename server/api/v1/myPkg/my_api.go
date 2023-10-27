package myPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type MyApi struct {
}

func (m *MyApi) CreateApi(c *gin.Context) {
	apiService.CreateApiS()
	response.Ok(c)
	//response.OkWithMessage("hello", c)
}

var (
	apiService = service.ServiceGroupApp.MyPkgServiceGroup
)
