import service from '@/utils/request'

// @Tags QmUser
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QmUser true "创建用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /qmUser/createQmUser [post]
export const createQmUser = (data) => {
  return service({
    url: '/qmUser/createQmUser',
    method: 'post',
    data
  })
}

// @Tags QmUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QmUser true "删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qmUser/deleteQmUser [delete]
export const deleteQmUser = (params) => {
  return service({
    url: '/qmUser/deleteQmUser',
    method: 'delete',
    params
  })
}

// @Tags QmUser
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /qmUser/deleteQmUser [delete]
export const deleteQmUserByIds = (params) => {
  return service({
    url: '/qmUser/deleteQmUserByIds',
    method: 'delete',
    params
  })
}

// @Tags QmUser
// @Summary 更新用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.QmUser true "更新用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /qmUser/updateQmUser [put]
export const updateQmUser = (data) => {
  return service({
    url: '/qmUser/updateQmUser',
    method: 'put',
    data
  })
}

// @Tags QmUser
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.QmUser true "用id查询用户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /qmUser/findQmUser [get]
export const findQmUser = (params) => {
  return service({
    url: '/qmUser/findQmUser',
    method: 'get',
    params
  })
}

// @Tags QmUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /qmUser/getQmUserList [get]
export const getQmUserList = (params) => {
  return service({
    url: '/qmUser/getQmUserList',
    method: 'get',
    params
  })
}


export const adminChangePassword = (data) => {
  return service({
    url: '/qmUser/adminChangePassword',
    method: 'put',
    data
  })
}

