package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderDetail 商旅机票分销订单详情接口
// alitrip.btrip.flight.distribution.order.detail
//
// 商旅机票分销订单详情接口
func AlitripBtripFlightDistributionOrderDetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderDetailAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionOrderDetailAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionOrderDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
