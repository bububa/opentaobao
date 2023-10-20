package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangepay 商旅机票改签支付
// alitrip.btrip.flight.distribution.change.pay
//
// 改签订单支付
func Alitripbtripflightdistributionchangepay(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangepayAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangepayAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangepayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
