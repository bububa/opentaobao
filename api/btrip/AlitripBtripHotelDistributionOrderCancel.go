package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripHotelDistributionOrderCancel 商旅酒店API分销取消订单
// alitrip.btrip.hotel.distribution.order.cancel
//
// 商旅酒店API分销取消订单
func AlitripBtripHotelDistributionOrderCancel(clt *core.SDKClient, req *btrip.AlitripBtripHotelDistributionOrderCancelAPIRequest, session string) (*btrip.AlitripBtripHotelDistributionOrderCancelAPIResponse, error) {
	var resp btrip.AlitripBtripHotelDistributionOrderCancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
