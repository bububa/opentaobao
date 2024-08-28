package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderPay 商旅机票分销-订单支付
// alitrip.btrip.flight.distribution.order.pay
//
// 商旅机票分销订单支付
func AlitripBtripFlightDistributionOrderPay(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderPayAPIRequest, resp *btrip.AlitripBtripFlightDistributionOrderPayAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
