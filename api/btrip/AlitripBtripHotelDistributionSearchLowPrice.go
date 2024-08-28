package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchLowPrice 商旅酒店api分销-酒店最低价
// alitrip.btrip.hotel.distribution.search.low.price
//
// 商旅酒店api分销-酒店最低价
func AlitripBtripHotelDistributionSearchLowPrice(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchLowPriceAPIRequest, resp *btrip.AlitripBtripHotelDistributionSearchLowPriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
