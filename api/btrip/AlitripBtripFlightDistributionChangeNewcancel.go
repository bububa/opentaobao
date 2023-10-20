package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangenewcancel 商旅机票改签取消
// alitrip.btrip.flight.distribution.change.newcancel
//
// 商旅机票改签取消
func Alitripbtripflightdistributionchangenewcancel(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangenewcancelAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangenewcancelAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangenewcancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
