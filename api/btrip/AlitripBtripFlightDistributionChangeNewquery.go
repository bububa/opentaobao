package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangenewquery 改签航班询价V2
// alitrip.btrip.flight.distribution.change.newquery
//
// 商旅机票改签航班询价
func Alitripbtripflightdistributionchangenewquery(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangenewqueryAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangenewqueryAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangenewqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
