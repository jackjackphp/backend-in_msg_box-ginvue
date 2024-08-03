package service

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/qmclient/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/qmclient/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gofrs/uuid/v5"
)

var QmUser = new(qmUser)

type qmUser struct{}

// CreateQmUser 创建用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *qmUser) CreateQmUser(qmUser *model.QmUser) (err error) {
	// 查询username是否已经注册
	var count int64
	err = global.GVA_DB.Model(&model.QmUser{}).Where("username = ?", qmUser.Username).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户名已注册")
	}
	// 把密码调整为加密方式
	hashPwd := utils.BcryptHash(qmUser.Password)
	qmUser.Password = hashPwd
	// uuid创建
	qmUser.UUID, _ = uuid.NewV4()
	err = global.GVA_DB.Create(qmUser).Error
	return err
}

// DeleteQmUser 删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *qmUser) DeleteQmUser(ID string) (err error) {
	err = global.GVA_DB.Delete(&model.QmUser{}, "id = ?", ID).Error
	return err
}

// DeleteQmUserByIds 批量删除用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *qmUser) DeleteQmUserByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]model.QmUser{}, "id in ?", IDs).Error
	return err
}

// UpdateQmUser 更新用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *qmUser) UpdateQmUser(qmUser model.QmUser) (err error) {
	err = global.GVA_DB.Model(&model.QmUser{}).Where("id = ?", qmUser.ID).Updates(&qmUser).Error
	return err
}

// GetQmUser 根据ID获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *qmUser) GetQmUser(ID string) (qmUser model.QmUser, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&qmUser).Error
	return
}

// GetQmUserInfoList 分页获取用户记录
// Author [piexlmax](https://github.com/piexlmax)
func (s *qmUser) GetQmUserInfoList(info request.QmUserSearch) (list []model.QmUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.QmUser{})
	var qmUsers []model.QmUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Username != "" {
		db = db.Where("username LIKE ?", "%"+info.Username+"%")
	}
	if info.Nickname != "" {
		db = db.Where("nickname LIKE ?", "%"+info.Nickname+"%")
	}
	if info.Gender != "" {
		db = db.Where("gender = ?", info.Gender)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&qmUsers).Error
	return qmUsers, total, err
}

// AdminChangePassword 请实现方法
// Author [yourname](https://github.com/yourname)
func (s *qmUser) AdminChangePassword(req request.AdminChangePasswordReq) (err error) {
	newPwd := utils.BcryptHash(req.Password)
	db := global.GVA_DB.Model(&model.QmUser{}).Where("id = ?", req.UserID).Update("password", newPwd)
	return db.Error
}
