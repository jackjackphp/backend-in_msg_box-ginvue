package response

import "github.com/eyotang/gin-vue-admin/server/model"

type SysMenusResponse struct {
	Menus []model.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []model.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu model.SysBaseMenu `json:"menu"`
}
