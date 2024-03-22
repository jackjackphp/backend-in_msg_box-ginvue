package webcms

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/webcms"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WebconfigApi struct{}

var (
	webconfigService = service.ServiceGroupApp.WebcmsServiceGroup.WebconfigService
)

func (w *WebconfigApi) CreateWebconfig(c *gin.Context) {
	var webconfig webcms.Webconfig
	err := c.ShouldBindJSON(&webconfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := webconfigService.CreateWebconfig(webconfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (w *WebconfigApi) GetWebconfig(c *gin.Context) {
	siteinfo := c.GetStringMapString("siteinfo")

	if list, err := webconfigService.GetWebconfig(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{
			"list":   list,
			"siteid": siteinfo["id"],
		}, "获取成功", c)
	}
}

func (w *WebconfigApi) DeleteWebconfig(c *gin.Context) {
	var id uint
	err := c.ShouldBindJSON(&id)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := webconfigService.DeleteWebconfig(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (w *WebconfigApi) SetWebconfig(c *gin.Context) {
	var webconfig webcms.Webconfig
	err := c.ShouldBindJSON(&webconfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := webconfigService.UpdatesWebconfig(webconfig); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.Ok(c)
	}
}

// 切换站点ChangeWebconfig
func (w *WebconfigApi) ChangeWebconfig(c *gin.Context) {
	var siteid int
	err := c.ShouldBindJSON(&siteid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_CONFIG.Website.Siteid = siteid

	response.Ok(c)

}