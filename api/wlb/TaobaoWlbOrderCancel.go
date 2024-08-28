package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbOrderCancel 取消物流宝订单
// taobao.wlb.order.cancel
//
// 取消物流宝订单
func TaobaoWlbOrderCancel(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbOrderCancelAPIRequest, resp *wlb.TaobaoWlbOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
