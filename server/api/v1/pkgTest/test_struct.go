package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TestStructApi struct {
}

var tsService = service.ServiceGroupApp.PkgTestServiceGroup.TestStructService

// CreateTestStruct 创建测试结构
// @Tags TestStruct
// @Summary 创建测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.TestStruct true "创建测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ts/createTestStruct [post]
func (tsApi *TestStructApi) CreateTestStruct(c *gin.Context) {
	var ts pkgTest.TestStruct
	err := c.ShouldBindJSON(&ts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tsService.CreateTestStruct(&ts); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTestStruct 删除测试结构
// @Tags TestStruct
// @Summary 删除测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.TestStruct true "删除测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ts/deleteTestStruct [delete]
func (tsApi *TestStructApi) DeleteTestStruct(c *gin.Context) {
	var ts pkgTest.TestStruct
	err := c.ShouldBindJSON(&ts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tsService.DeleteTestStruct(ts); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTestStructByIds 批量删除测试结构
// @Tags TestStruct
// @Summary 批量删除测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ts/deleteTestStructByIds [delete]
func (tsApi *TestStructApi) DeleteTestStructByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tsService.DeleteTestStructByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTestStruct 更新测试结构
// @Tags TestStruct
// @Summary 更新测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.TestStruct true "更新测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ts/updateTestStruct [put]
func (tsApi *TestStructApi) UpdateTestStruct(c *gin.Context) {
	var ts pkgTest.TestStruct
	err := c.ShouldBindJSON(&ts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tsService.UpdateTestStruct(ts); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTestStruct 用id查询测试结构
// @Tags TestStruct
// @Summary 用id查询测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest.TestStruct true "用id查询测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ts/findTestStruct [get]
func (tsApi *TestStructApi) FindTestStruct(c *gin.Context) {
	var ts pkgTest.TestStruct
	err := c.ShouldBindQuery(&ts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rets, err := tsService.GetTestStruct(ts.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rets": rets}, c)
	}
}

// GetTestStructList 分页获取测试结构列表
// @Tags TestStruct
// @Summary 分页获取测试结构列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTestReq.TestStructSearch true "分页获取测试结构列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ts/getTestStructList [get]
func (tsApi *TestStructApi) GetTestStructList(c *gin.Context) {
	var pageInfo pkgTestReq.TestStructSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tsService.GetTestStructInfoList(pageInfo); err != nil {
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
