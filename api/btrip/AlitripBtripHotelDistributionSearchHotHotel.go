package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionsearchhothotel 商旅酒店api分销-热点酒店
// alitrip.btrip.hotel.distribution.search.hot.hotel
//
// 商旅酒店api分销-热点酒店
func Alitripbtriphoteldistributionsearchhothotel(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionsearchhothotelAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionsearchhothotelAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionsearchhothotelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
