package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionrefundnewdetail 商旅机票退票详情接口V2
// alitrip.btrip.flight.distribution.refund.newdetail
//
// 商旅机票退票详情接口V2
func Alitripbtripflightdistributionrefundnewdetail(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionrefundnewdetailAPIRequest, session string) (*btrip.AlitripbtripflightdistributionrefundnewdetailAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionrefundnewdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
