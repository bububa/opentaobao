package tvupadmin

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvupadmin"
)

// YunosTvpubadminUserOrderlist 获取用户订单列表
// yunos.tvpubadmin.user.orderlist
//
// 获取用户订单列表
func YunosTvpubadminUserOrderlist(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminUserOrderlistAPIRequest, resp *tvupadmin.YunosTvpubadminUserOrderlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
