package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewpay 商旅机票改签支付V2
// alitrip.btrip.flight.distribution.change.newpay
//
// 商旅机票改签支付V2
func AlitripBtripFlightDistributionChangeNewpay(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewpayAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeNewpayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
