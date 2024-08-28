package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderCreate 商旅机票分销-创建订单
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
func AlitripBtripFlightDistributionOrderCreate(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderCreateAPIRequest, resp *btrip.AlitripBtripFlightDistributionOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
