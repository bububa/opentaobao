package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionmodifyflightsearch 改签航班列表
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
func Alitripbtripflightdistributionmodifyflightsearch(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionmodifyflightsearchAPIRequest, session string) (*btrip.AlitripbtripflightdistributionmodifyflightsearchAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionmodifyflightsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
