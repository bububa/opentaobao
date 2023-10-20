package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchHotHotel 商旅酒店api分销-热点酒店
// alitrip.btrip.hotel.distribution.search.hot.hotel
//
// 商旅酒店api分销-热点酒店
func AlitripBtripHotelDistributionSearchHotHotel(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchHotHotelAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionSearchHotHotelAPIResponse, error) {
	var resp btrip.AlitripBtripHotelDistributionSearchHotHotelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
