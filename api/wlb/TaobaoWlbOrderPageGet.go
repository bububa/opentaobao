package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderPageGet 分页查询物流宝订单
// taobao.wlb.order.page.get
//
// 分页查询物流宝订单
func TaobaoWlbOrderPageGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbOrderPageGetAPIRequest, resp *wlb.TaobaoWlbOrderPageGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
