package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionSearchDetail 商旅酒店api分销-详情报价接口
// alitrip.btrip.hotel.distribution.search.detail
//
// 商旅酒店api分销-详情报价接口
func AlitripBtripHotelDistributionSearchDetail(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionSearchDetailAPIRequest, resp *btrip.AlitripBtripHotelDistributionSearchDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
