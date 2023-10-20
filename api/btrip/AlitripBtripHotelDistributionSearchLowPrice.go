package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchLowPrice 商旅酒店api分销-酒店最低价
// alitrip.btrip.hotel.distribution.search.low.price
//
// 商旅酒店api分销-酒店最低价
func AlitripBtripHotelDistributionSearchLowPrice(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchLowPriceAPIRequest, resp *btrip.AlitripBtripHotelDistributionSearchLowPriceAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
