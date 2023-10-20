package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderCreate 商旅机票分销-创建订单
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
func AlitripBtripFlightDistributionOrderCreate(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderCreateAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionOrderCreateAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionOrderCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
