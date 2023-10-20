package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtriphoteldistributionsearchlowprice 商旅酒店api分销-酒店最低价
// alitrip.btrip.hotel.distribution.search.low.price
//
// 商旅酒店api分销-酒店最低价
func Alitripbtriphoteldistributionsearchlowprice(clt *core.SDKClient, req *btrip.AlitripbtriphoteldistributionsearchlowpriceAPIRequest, session string) (*btrip.AlitripbtriphoteldistributionsearchlowpriceAPIResponse, error) {
	var resp btrip.AlitripbtriphoteldistributionsearchlowpriceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
