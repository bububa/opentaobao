package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangecancel 商旅机票改签取消
// alitrip.btrip.flight.distribution.change.cancel
//
// 商旅机票改签取消
func Alitripbtripflightdistributionchangecancel(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangecancelAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangecancelAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangecancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
