package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderDetail 商旅酒店API分销查询订单详情
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
func AlitripBtripHotelDistributionOrderDetail(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderDetailAPIRequest, resp *btrip.AlitripBtripHotelDistributionOrderDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
