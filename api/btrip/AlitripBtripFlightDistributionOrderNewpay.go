package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderNewpay 商旅机票分销-订单支付V2
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
func AlitripBtripFlightDistributionOrderNewpay(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderNewpayAPIRequest, resp *btrip.AlitripBtripFlightDistributionOrderNewpayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
