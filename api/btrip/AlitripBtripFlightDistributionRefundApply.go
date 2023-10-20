package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionrefundapply 商旅机票分销-退票申请
// alitrip.btrip.flight.distribution.refund.apply
//
// 商旅机票分销-退票申请
func Alitripbtripflightdistributionrefundapply(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionrefundapplyAPIRequest, session string) (*btrip.AlitripbtripflightdistributionrefundapplyAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionrefundapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
