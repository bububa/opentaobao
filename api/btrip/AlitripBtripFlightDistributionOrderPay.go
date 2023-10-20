package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionorderpay 商旅机票分销-订单支付
// alitrip.btrip.flight.distribution.order.pay
//
// 商旅机票分销订单支付
func Alitripbtripflightdistributionorderpay(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionorderpayAPIRequest, session string) (*btrip.AlitripbtripflightdistributionorderpayAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionorderpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
