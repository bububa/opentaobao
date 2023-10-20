package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionmodifynewflightsearch 改签航班列表V2
// alitrip.btrip.flight.distribution.modify.newflightsearch
//
// 改签航班列表V2
func Alitripbtripflightdistributionmodifynewflightsearch(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionmodifynewflightsearchAPIRequest, session string) (*btrip.AlitripbtripflightdistributionmodifynewflightsearchAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionmodifynewflightsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
