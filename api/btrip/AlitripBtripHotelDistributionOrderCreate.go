package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderCreate 商旅酒店分销-创建订单
// alitrip.btrip.hotel.distribution.order.create
//
// 商旅酒店分销-创建订单
func AlitripBtripHotelDistributionOrderCreate(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderCreateAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionOrderCreateAPIResponse, error) {
	var resp btrip.AlitripBtripHotelDistributionOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
