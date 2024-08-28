package ottpay

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ottpay"
)

// YoukuOttPayOrderDeleteorder 退订应用中心支付订单
// youku.ott.pay.order.deleteorder
//
// 应用中心sdk连续包月退订接口
func YoukuOttPayOrderDeleteorder(ctx context.Context, clt *core.SDKClient, req *ottpay.YoukuOttPayOrderDeleteorderAPIRequest, resp *ottpay.YoukuOttPayOrderDeleteorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
