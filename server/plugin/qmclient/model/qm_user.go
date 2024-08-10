package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gofrs/uuid/v5"
)

var _ system.Login = (*QmUser)(nil)

// QmUser 用户 结构体
type QmUser struct {
	global.GVA_MODEL
	Username    string    `json:"username" form:"username" gorm:"column:username;comment:用户名;" binding:"required"`                              //用户名
	Password    string    `json:"password" form:"password" gorm:"column:password;comment:密码;" binding:"required"`                               //密码
	Avatar      string    `json:"avatar" form:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;column:avatar;comment:头像;"` //头像
	Nickname    string    `json:"nickname" form:"nickname" gorm:"default:新用户;column:nickname;comment:昵称;"`                                      //昵称
	UUID        uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:UUID;"`                                                            //UUID
	Gender      string    `json:"gender" form:"gender" gorm:"column:gender;comment:性别;"`                                                        //性别
	Desc        string    `json:"desc" form:"desc" gorm:"column:desc;comment:简介;type:text;"`                                                    //简介
	AuthorityId uint      `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:角色ID"`                                       //角色ID
}

func (u QmUser) GetUsername() string {
	return u.Username
}

func (u QmUser) GetNickname() string {
	return u.Nickname
}

func (u QmUser) GetUUID() uuid.UUID {
	return u.UUID
}

func (u QmUser) GetUserId() uint {
	return u.ID
}

func (u QmUser) GetAuthorityId() uint {
	return u.AuthorityId
}

func (u QmUser) GetUserInfo() any {
	return u
}

// TableName 用户 QmUser自定义表名 qm_client_user
func (QmUser) TableName() string {
	return "qm_client_user"
}
