import service from '@/utils/request'

// @Tags TestStruct
// @Summary 创建测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct true "创建测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /ts/createTestStruct [post]
export const createTestStruct = (data) => {
  return service({
    url: '/ts/createTestStruct',
    method: 'post',
    data
  })
}

// @Tags TestStruct
// @Summary 删除测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct true "删除测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ts/deleteTestStruct [delete]
export const deleteTestStruct = (data) => {
  return service({
    url: '/ts/deleteTestStruct',
    method: 'delete',
    data
  })
}

// @Tags TestStruct
// @Summary 批量删除测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ts/deleteTestStruct [delete]
export const deleteTestStructByIds = (data) => {
  return service({
    url: '/ts/deleteTestStructByIds',
    method: 'delete',
    data
  })
}

// @Tags TestStruct
// @Summary 更新测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStruct true "更新测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ts/updateTestStruct [put]
export const updateTestStruct = (data) => {
  return service({
    url: '/ts/updateTestStruct',
    method: 'put',
    data
  })
}

// @Tags TestStruct
// @Summary 用id查询测试结构
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestStruct true "用id查询测试结构"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ts/findTestStruct [get]
export const findTestStruct = (params) => {
  return service({
    url: '/ts/findTestStruct',
    method: 'get',
    params
  })
}

// @Tags TestStruct
// @Summary 分页获取测试结构列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试结构列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ts/getTestStructList [get]
export const getTestStructList = (params) => {
  return service({
    url: '/ts/getTestStructList',
    method: 'get',
    params
  })
}
