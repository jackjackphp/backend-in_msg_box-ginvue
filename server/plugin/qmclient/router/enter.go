package router

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/qmclient/api"

var (
	Router    = new(router)
	apiQmUser = api.Api.QmUser
)

type router struct{ QmUser qmUser }
