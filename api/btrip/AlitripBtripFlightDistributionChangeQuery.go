package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangequery 改签航班询价
// alitrip.btrip.flight.distribution.change.query
//
// 商旅机票改签航班询价
func Alitripbtripflightdistributionchangequery(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangequeryAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangequeryAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
