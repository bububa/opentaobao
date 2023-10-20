package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionrefundprecal 商旅机票分销-退票费预计算
// alitrip.btrip.flight.distribution.refund.precal
//
// 商旅机票分销-退票费预计算
func Alitripbtripflightdistributionrefundprecal(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionrefundprecalAPIRequest, session string) (*btrip.AlitripbtripflightdistributionrefundprecalAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionrefundprecalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
