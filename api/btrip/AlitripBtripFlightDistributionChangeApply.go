package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeApply 商旅机票改签申请
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
func AlitripBtripFlightDistributionChangeApply(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeApplyAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeApplyAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
