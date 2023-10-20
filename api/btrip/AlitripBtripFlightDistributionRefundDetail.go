package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionrefunddetail 商旅机票退票详情接口
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
func Alitripbtripflightdistributionrefunddetail(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionrefunddetailAPIRequest, session string) (*btrip.AlitripbtripflightdistributionrefunddetailAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionrefunddetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
