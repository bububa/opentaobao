package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchHotHotel 商旅酒店api分销-热点酒店
// alitrip.btrip.hotel.distribution.search.hot.hotel
//
// 商旅酒店api分销-热点酒店
func AlitripBtripHotelDistributionSearchHotHotel(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchHotHotelAPIRequest, resp *btrip.AlitripBtripHotelDistributionSearchHotHotelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
