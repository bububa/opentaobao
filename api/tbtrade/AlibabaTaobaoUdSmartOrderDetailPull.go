package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// AlibabaTaobaoUdSmartOrderDetailPull UD效果外投订单明细拉取
// alibaba.taobao.ud.smart.order.detail.pull
//
// UD效果外投订单明细拉取
func AlibabaTaobaoUdSmartOrderDetailPull(ctx context.Context, clt *core.SDKClient, req *tbtrade.AlibabaTaobaoUdSmartOrderDetailPullAPIRequest, resp *tbtrade.AlibabaTaobaoUdSmartOrderDetailPullAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
