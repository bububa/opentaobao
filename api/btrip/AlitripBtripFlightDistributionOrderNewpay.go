package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderNewpay 商旅机票分销-订单支付V2
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
func AlitripBtripFlightDistributionOrderNewpay(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderNewpayAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionOrderNewpayAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionOrderNewpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
