package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderCreate 商旅酒店分销-创建订单
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
func AlitripBtripHotelDistributionOrderCreate(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderCreateAPIRequest, resp *btrip.AlitripBtripHotelDistributionOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
