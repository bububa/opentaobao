package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderDetail 商旅酒店API分销查询订单详情
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
func AlitripBtripHotelDistributionOrderDetail(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderDetailAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionOrderDetailAPIResponse, error) {
	var resp btrip.AlitripBtripHotelDistributionOrderDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
