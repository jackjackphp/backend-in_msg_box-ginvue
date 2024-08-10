package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/qmclient/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/qmclient/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/qmclient/plugin"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var QmUser = new(qmUser)

type qmUser struct{}

// CreateQmUser 创建用户
// @Tags QmUser
// @Summary 创建用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body qmclient.QmUser true "创建用户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /qmUser/createQmUser [post]
func (a *qmUser) CreateQmUser(c *gin.Context) {
	var info model.QmUser
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serviceQmUser.CreateQmUser(&info)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteQmUser 删除用户
// @Tags QmUser
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body qmclient.QmUser true "删除用户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /qmUser/deleteQmUser [delete]
func (a *qmUser) DeleteQmUser(c *gin.Context) {
	ID := c.Query("ID")
	err := serviceQmUser.DeleteQmUser(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteQmUserByIds 批量删除用户
// @Tags QmUser
// @Summary 批量删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /qmUser/deleteQmUserByIds [delete]
func (a *qmUser) DeleteQmUserByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := serviceQmUser.DeleteQmUserByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateQmUser 更新用户
// @Tags QmUser
// @Summary 更新用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body qmclient.QmUser true "更新用户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /qmUser/updateQmUser [put]
func (a *qmUser) UpdateQmUser(c *gin.Context) {
	var info model.QmUser
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info.Username = ""
	info.Password = ""
	err = serviceQmUser.UpdateQmUser(info)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindQmUser 用id查询用户
// @Tags QmUser
// @Summary 用id查询用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query qmclient.QmUser true "用id查询用户"
// @Success 200 {object} response.Response{data=object{reqmUser=qmclient.QmUser},msg=string} "查询成功"
// @Router /qmUser/findQmUser [get]
func (a *qmUser) FindQmUser(c *gin.Context) {
	ID := c.Query("ID")
	reqmUser, err := serviceQmUser.GetQmUser(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(reqmUser, c)
}

// GetQmUserList 分页获取用户列表
// @Tags QmUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.QmUserSearch true "分页获取用户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /qmUser/getQmUserList [get]
func (a *qmUser) GetQmUserList(c *gin.Context) {
	var pageInfo request.QmUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serviceQmUser.GetQmUserInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetQmUserPublic 不需要鉴权的用户接口
// @Tags QmUser
// @Summary 不需要鉴权的用户接口
// @accept application/json
// @Produce application/json
// @Param data query request.QmUserSearch true "分页获取用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/getQmUserPublic [get]
func (a *qmUser) GetQmUserPublic(c *gin.Context) {
	// 此接口不需要鉴权 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{"info": "不需要鉴权的用户接口信息"}, "获取成功", c)
}

// AdminChangePassword 等待开发的的用户接口
// @Tags QmUser
// @Summary 等待开发的的用户接口
// @accept application/json
// @Produce application/json
// @Param data body request.AdminChangePasswordReq true "分页获取用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/adminChangePassword [PUT]
func (a *qmUser) AdminChangePassword(c *gin.Context) {
	var req request.AdminChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 请添加自己的业务逻辑
	err = serviceQmUser.AdminChangePassword(req)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData("返回数据", c)
}

// Register 等待开发的的用户接口
// @Tags QmUser
// @Summary 等待开发的的用户接口
// @accept application/json
// @Produce application/json
// @Param data query request.QmUserSearch true "分页获取用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/register [POST]
func (a *qmUser) Register(c *gin.Context) {
	var registerReq request.QmUser
	err := c.ShouldBindJSON(&registerReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var user model.QmUser
	user.Username = registerReq.Username
	user.Password = registerReq.Password
	user.Nickname = registerReq.Nickname
	user.AuthorityId = plugin.Config.AuthorityID

	// 请添加自己的业务逻辑
	err = serviceQmUser.CreateQmUser(&user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithMessage("注册失败:"+err.Error(), c)
		return
	}
	response.OkWithData("注册成功", c)
}

// Login 等待开发的的用户接口
// @Tags QmUser
// @Summary 等待开发的的用户接口
// @accept application/json
// @Produce application/json
// @Param data query request.QmUserSearch true "分页获取用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/login [POST]
func (a *qmUser) Login(c *gin.Context) {
	var loginReq request.QmUser
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 请添加自己的业务逻辑
	user, err := serviceQmUser.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	token, claims, err := utils.LoginToken(user)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}

	response.OkWithData(gin.H{
		"token":     token,
		"expiresAt": claims.RegisteredClaims.ExpiresAt,
		"user":      user,
	}, c)
}

// GetUserInfo 等待开发的的用户接口
// @Tags QmUser
// @Summary 等待开发的的用户接口
// @accept application/json
// @Produce application/json
// @Param data query request.QmUserSearch true "分页获取用户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /qmUser/getUserInfo [GET]
func (a *qmUser) GetUserInfo(c *gin.Context) {
	// 请添加自己的业务逻辑
	id := utils.GetUserID(c)
	if id == 0 {
		response.FailWithMessage("获取失败", c)
		return
	}
	user, err := serviceQmUser.GetUserInfo(id)
	if err != nil {
		global.GVA_LOG.Error("失败!", zap.Error(err))
		response.FailWithMessage("失败", c)
		return
	}
	response.OkWithData(user, c)
}
