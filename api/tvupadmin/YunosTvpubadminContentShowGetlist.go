package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowGetlist 节目审核获取节目列表
// yunos.tvpubadmin.content.show.getlist
//
// 节目审核获取节目列表
func YunosTvpubadminContentShowGetlist(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowGetlistAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowGetlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
