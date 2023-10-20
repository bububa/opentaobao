package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangedetail 商旅机票改签详情接口
// alitrip.btrip.flight.distribution.change.detail
//
// 商旅机票改签详情接口
func Alitripbtripflightdistributionchangedetail(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangedetailAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangedetailAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
