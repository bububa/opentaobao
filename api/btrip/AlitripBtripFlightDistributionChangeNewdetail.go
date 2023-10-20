package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionchangenewdetail 商旅机票改签详情接口
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
func Alitripbtripflightdistributionchangenewdetail(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionchangenewdetailAPIRequest, session string) (*btrip.AlitripbtripflightdistributionchangenewdetailAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionchangenewdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
