package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionflightlist 商旅机票航班列表接口
// alitrip.btrip.flight.distribution.flightlist
//
// 商旅机票航班列表接口，用于分销询价
func Alitripbtripflightdistributionflightlist(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionflightlistAPIRequest, session string) (*btrip.AlitripbtripflightdistributionflightlistAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionflightlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
