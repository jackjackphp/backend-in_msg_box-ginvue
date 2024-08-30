package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Api(ctx context.Context) {
	entities := []model.SysApi{{Path: "/qmUser/createQmUser", Description: "新增用户", ApiGroup: "用户", Method: "POST"}, {Path: "/qmUser/deleteQmUser", Description: "删除用户", ApiGroup: "用户", Method: "DELETE"}, {Path: "/qmUser/deleteQmUserByIds", Description: "批量删除用户", ApiGroup: "用户", Method: "DELETE"}, {Path: "/qmUser/updateQmUser", Description: "更新用户", ApiGroup: "用户", Method: "PUT"}, {Path: "/qmUser/findQmUser", Description: "根据ID获取用户", ApiGroup: "用户", Method: "GET"}, {Path: "/qmUser/getQmUserList", Description: "获取用户列表", ApiGroup: "用户", Method: "GET"}, {Path: "/qmUser/adminChangePassword", Description: "修改用户密码", ApiGroup: "用户", Method: "PUT"}, {Path: "/qmUser/getUserInfo", Description: "获取用户信息", ApiGroup: "用户", Method: "GET"}}
	utils.RegisterApis(entities...)
}
