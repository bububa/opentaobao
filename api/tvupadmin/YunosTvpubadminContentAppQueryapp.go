package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAppQueryapp 查询应用信息
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
func YunosTvpubadminContentAppQueryapp(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAppQueryappAPIRequest, resp *tvupadmin.YunosTvpubadminContentAppQueryappAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
