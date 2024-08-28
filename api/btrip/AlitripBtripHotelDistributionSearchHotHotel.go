package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchHotHotel 商旅酒店api分销-热点酒店
// alitrip.btrip.hotel.distribution.search.hot.hotel
//
// 商旅酒店api分销-热点酒店
func AlitripBtripHotelDistributionSearchHotHotel(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchHotHotelAPIRequest, resp *btrip.AlitripBtripHotelDistributionSearchHotHotelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
