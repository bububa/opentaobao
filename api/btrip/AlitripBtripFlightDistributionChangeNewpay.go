package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewpay 商旅机票改签支付V2
// alitrip.btrip.flight.distribution.change.newpay
//
// 商旅机票改签支付V2
func AlitripBtripFlightDistributionChangeNewpay(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewpayAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeNewpayAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeNewpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
