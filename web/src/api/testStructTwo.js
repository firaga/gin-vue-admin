import service from '@/utils/request'

// @Tags TestStructTwo
// @Summary 创建测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStructTwo true "创建测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tst/createTestStructTwo [post]
export const createTestStructTwo = (data) => {
  return service({
    url: '/tst/createTestStructTwo',
    method: 'post',
    data
  })
}

// @Tags TestStructTwo
// @Summary 删除测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStructTwo true "删除测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tst/deleteTestStructTwo [delete]
export const deleteTestStructTwo = (data) => {
  return service({
    url: '/tst/deleteTestStructTwo',
    method: 'delete',
    data
  })
}

// @Tags TestStructTwo
// @Summary 批量删除测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tst/deleteTestStructTwo [delete]
export const deleteTestStructTwoByIds = (data) => {
  return service({
    url: '/tst/deleteTestStructTwoByIds',
    method: 'delete',
    data
  })
}

// @Tags TestStructTwo
// @Summary 更新测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TestStructTwo true "更新测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tst/updateTestStructTwo [put]
export const updateTestStructTwo = (data) => {
  return service({
    url: '/tst/updateTestStructTwo',
    method: 'put',
    data
  })
}

// @Tags TestStructTwo
// @Summary 用id查询测试结构二
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TestStructTwo true "用id查询测试结构二"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tst/findTestStructTwo [get]
export const findTestStructTwo = (params) => {
  return service({
    url: '/tst/findTestStructTwo',
    method: 'get',
    params
  })
}

// @Tags TestStructTwo
// @Summary 分页获取测试结构二列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取测试结构二列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tst/getTestStructTwoList [get]
export const getTestStructTwoList = (params) => {
  return service({
    url: '/tst/getTestStructTwoList',
    method: 'get',
    params
  })
}
