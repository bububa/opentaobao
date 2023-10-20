package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangeapply 商旅机票改签申请
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
func Alitripbtripflightdistributionchangeapply(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangeapplyAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangeapplyAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangeapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
