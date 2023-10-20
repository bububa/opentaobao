package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionordernewpay 商旅机票分销-订单支付V2
// alitrip.btrip.flight.distribution.order.newpay
//
// 商旅机票分销-订单支付V2
func Alitripbtripflightdistributionordernewpay(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionordernewpayAPIRequest, session string) (*btrip.AlitripbtripflightdistributionordernewpayAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionordernewpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
