package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderCreate 商旅酒店分销-创建订单
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
func AlitripBtripHotelDistributionOrderCreate(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderCreateAPIRequest, resp *btrip.AlitripBtripHotelDistributionOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
