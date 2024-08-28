package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderCancel 商旅酒店API分销取消订单
// alitrip.btrip.hotel.distribution.order.cancel
//
// 商旅酒店API分销取消订单
func AlitripBtripHotelDistributionOrderCancel(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderCancelAPIRequest, resp *btrip.AlitripBtripHotelDistributionOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
