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
