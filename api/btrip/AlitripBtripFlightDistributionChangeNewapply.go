package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangenewapply 商旅机票改签申请V2
// alitrip.btrip.flight.distribution.change.newapply
//
// 商旅机票改签申请
func Alitripbtripflightdistributionchangenewapply(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangenewapplyAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangenewapplyAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangenewapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
