package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionorderdetail 商旅机票分销订单详情接口
// alitrip.btrip.flight.distribution.order.detail
//
// 商旅机票分销订单详情接口
func Alitripbtripflightdistributionorderdetail(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionorderdetailAPIRequest, session string) (*btrip.AlitripbtripflightdistributionorderdetailAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionorderdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
