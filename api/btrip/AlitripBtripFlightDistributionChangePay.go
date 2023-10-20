package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangePay 商旅机票改签支付
// alitrip.btrip.flight.distribution.change.pay
//
// 改签订单支付
func AlitripBtripFlightDistributionChangePay(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangePayAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangePayAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangePayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
