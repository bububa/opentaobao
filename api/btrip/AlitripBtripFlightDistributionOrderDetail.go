package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderDetail 商旅机票分销订单详情接口
// alitrip.btrip.flight.distribution.order.detail
//
// 商旅机票分销订单详情接口
func AlitripBtripFlightDistributionOrderDetail(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderDetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionOrderDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
