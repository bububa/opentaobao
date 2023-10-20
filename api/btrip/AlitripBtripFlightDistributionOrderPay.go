package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderPay 商旅机票分销-订单支付
// alitrip.btrip.flight.distribution.order.pay
//
// 商旅机票分销订单支付
func AlitripBtripFlightDistributionOrderPay(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderPayAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionOrderPayAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionOrderPayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
