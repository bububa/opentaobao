package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderCreate 创建物流宝订单
// taobao.wlb.order.create
//
// 创建物流宝订单，由外部ISV或者ERP，Elink，淘宝交易产生
func TaobaoWlbOrderCreate(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbOrderCreateAPIRequest, resp *wlb.TaobaoWlbOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
