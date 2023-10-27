package pkgTest2

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2"
	pkgTest2Req "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest2/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestStructTwoApi struct {
}

var tstService = service.ServiceGroupApp.PkgTest2ServiceGroup.TestStructTwoService

// CreateTestStructTwo 创建测试结构二
// @Tags TestStructTwo
// @Summary 创建测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest2.TestStructTwo true "创建测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tst/createTestStructTwo [post]
func (tstApi *TestStructTwoApi) CreateTestStructTwo(c *gin.Context) {
	var tst pkgTest2.TestStructTwo
	err := c.ShouldBindJSON(&tst)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tstService.CreateTestStructTwo(&tst); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStructTwo 删除测试结构二
// @Tags TestStructTwo
// @Summary 删除测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest2.TestStructTwo true "删除测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tst/deleteTestStructTwo [delete]
func (tstApi *TestStructTwoApi) DeleteTestStructTwo(c *gin.Context) {
	var tst pkgTest2.TestStructTwo
	err := c.ShouldBindJSON(&tst)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tstService.DeleteTestStructTwo(tst); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStructTwoByIds 批量删除测试结构二
// @Tags TestStructTwo
// @Summary 批量删除测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tst/deleteTestStructTwoByIds [delete]
func (tstApi *TestStructTwoApi) DeleteTestStructTwoByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tstService.DeleteTestStructTwoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStructTwo 更新测试结构二
// @Tags TestStructTwo
// @Summary 更新测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest2.TestStructTwo true "更新测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tst/updateTestStructTwo [put]
func (tstApi *TestStructTwoApi) UpdateTestStructTwo(c *gin.Context) {
	var tst pkgTest2.TestStructTwo
	err := c.ShouldBindJSON(&tst)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tstService.UpdateTestStructTwo(tst); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStructTwo 用id查询测试结构二
// @Tags TestStructTwo
// @Summary 用id查询测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest2.TestStructTwo true "用id查询测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tst/findTestStructTwo [get]
func (tstApi *TestStructTwoApi) FindTestStructTwo(c *gin.Context) {
	var tst pkgTest2.TestStructTwo
	err := c.ShouldBindQuery(&tst)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retst, err := tstService.GetTestStructTwo(tst.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retst": retst}, c)
	}
}

// GetTestStructTwoList 分页获取测试结构二列表
// @Tags TestStructTwo
// @Summary 分页获取测试结构二列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest2Req.TestStructTwoSearch true "分页获取测试结构二列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tst/getTestStructTwoList [get]
func (tstApi *TestStructTwoApi) GetTestStructTwoList(c *gin.Context) {
	var pageInfo pkgTest2Req.TestStructTwoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tstService.GetTestStructTwoInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
