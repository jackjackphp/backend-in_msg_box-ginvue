package initialize

import (
	"context"
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

func Menu(ctx context.Context) {
	entities := []model.SysBaseMenu{{ParentId: 0, Path: "qmclientMenu", Name: "qmclientMenu", Hidden: false, Component: "view/routerHolder.vue", Sort: 0, Meta: model.Meta{Title: "用户管理", Icon: "school"}}, {ParentId: 0, Path: "qmUser", Name: "qmUser", Hidden: false, Component: "plugin/qmclient/view/qmUser.vue", Sort: 0, Meta: model.Meta{Title: "用户", Icon: ""}}}
	utils.RegisterMenus(entities...)
}
