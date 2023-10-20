package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAppQuerybylicence 按牌照查询应用
// yunos.tvpubadmin.content.app.querybylicence
//
// 按牌照查询应用
func YunosTvpubadminContentAppQuerybylicence(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAppQuerybylicenceAPIRequest, resp *tvupadmin.YunosTvpubadminContentAppQuerybylicenceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
