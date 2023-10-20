package tvupadmin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminUserOrderlist 获取用户订单列表
// yunos.tvpubadmin.user.orderlist
//
// 获取用户订单列表
func YunosTvpubadminUserOrderlist(clt *core.SDKClient, req *tvupadmin.YunosTvpubadminUserOrderlistAPIRequest, resp *tvupadmin.YunosTvpubadminUserOrderlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
