package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type QmUserSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Username       string     `json:"username" form:"username" `
	Nickname       string     `json:"nickname" form:"nickname" `
	Gender         string     `json:"gender" form:"gender" `
	request.PageInfo
}

type AdminChangePasswordReq struct {
	UserID   uint   `json:"userID"`
	Password string `json:"password"`
}

type QmUser struct {
	Username string `json:"username" form:"username" gorm:"column:username;comment:用户名;" binding:"required"` //用户名
	Password string `json:"password" form:"password" gorm:"column:password;comment:密码;" binding:"required"`  //密码
	Nickname string `json:"nickname" form:"nickname" gorm:"default:新用户;column:nickname;comment:昵称;"`         //昵称
}
