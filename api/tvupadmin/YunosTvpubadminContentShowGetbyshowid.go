package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentShowGetbyshowid 迎客松根据节目id获取节目元数据
// yunos.tvpubadmin.content.show.getbyshowid
//
// 迎客松根据节目id获取节目元数据
func YunosTvpubadminContentShowGetbyshowid(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentShowGetbyshowidAPIRequest, resp *tvupadmin.YunosTvpubadminContentShowGetbyshowidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
