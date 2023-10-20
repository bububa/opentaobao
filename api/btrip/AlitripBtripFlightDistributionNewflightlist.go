package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// Alitripbtripflightdistributionnewflightlist 商旅机票航班列表接口，用于分销询价V2
// alitrip.btrip.flight.distribution.newflightlist
//
// 商旅机票航班列表接口，用于分销询价V2
func Alitripbtripflightdistributionnewflightlist(clt *core.SDKClient, req *btrip.AlitripbtripflightdistributionnewflightlistAPIRequest, session string) (*btrip.AlitripbtripflightdistributionnewflightlistAPIResponse, error) {
	var resp btrip.AlitripbtripflightdistributionnewflightlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
