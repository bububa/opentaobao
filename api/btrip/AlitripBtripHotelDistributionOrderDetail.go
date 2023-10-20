package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderDetail 商旅酒店API分销查询订单详情
// alitrip.btrip.hotel.distribution.order.detail
//
// 商旅酒店API分销查询订单详情
func AlitripBtripHotelDistributionOrderDetail(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderDetailAPIRequest, resp *btrip.AlitripBtripHotelDistributionOrderDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
