package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeDetail 商旅机票改签详情接口
// alitrip.btrip.flight.distribution.change.detail
//
// 商旅机票改签详情接口
func AlitripBtripFlightDistributionChangeDetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeDetailAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeDetailAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
