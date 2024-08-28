package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderitemPageGet 分页查询物流宝订单商品详情
// taobao.wlb.orderitem.page.get
//
// 分页查询物流宝订单商品详情
func TaobaoWlbOrderitemPageGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbOrderitemPageGetAPIRequest, resp *wlb.TaobaoWlbOrderitemPageGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
