package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionordercancel 商旅机票分销-取消订单
// alitrip.btrip.flight.distribution.order.cancel
//
// 商旅机票分销取消订单
func Alitripbtripflightdistributionordercancel(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionordercancelAPIRequest, session string) (*btrip.AlitripbtripflightdistributionordercancelAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
