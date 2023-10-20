package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangenewpay 商旅机票改签支付V2
// alitrip.btrip.flight.distribution.change.newpay
//
// 商旅机票改签支付V2
func Alitripbtripflightdistributionchangenewpay(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangenewpayAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangenewpayAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangenewpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
