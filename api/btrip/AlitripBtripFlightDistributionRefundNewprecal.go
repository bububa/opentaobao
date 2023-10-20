package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionrefundnewprecal 商旅机票分销-退票费预计算
// alitrip.btrip.flight.distribution.refund.newprecal
//
// 商旅机票分销-退票费预计算
func Alitripbtripflightdistributionrefundnewprecal(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionrefundnewprecalAPIRequest, session string) (*btrip.AlitripbtripflightdistributionrefundnewprecalAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionrefundnewprecalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
