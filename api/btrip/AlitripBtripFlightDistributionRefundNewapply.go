package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionrefundnewapply 商旅机票分销-退票申请
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
func Alitripbtripflightdistributionrefundnewapply(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionrefundnewapplyAPIRequest, session string) (*btrip.AlitripbtripflightdistributionrefundnewapplyAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionrefundnewapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
