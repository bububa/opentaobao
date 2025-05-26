package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// AlibabaTaobaoUdSmartOrderCollectCartDetailPull UD效果外投收藏加购明细拉取
// alibaba.taobao.ud.smart.order.collect.cart.detail.pull
//
// UD效果外投收藏加购明细拉取
func AlibabaTaobaoUdSmartOrderCollectCartDetailPull(ctx context.Context, clt *core.SDKClient, req *tbtrade.AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIRequest, resp *tbtrade.AlibabaTaobaoUdSmartOrderCollectCartDetailPullAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
