package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewapply 商旅机票改签申请V2
// alitrip.btrip.flight.distribution.change.newapply
//
// 商旅机票改签申请
func AlitripBtripFlightDistributionChangeNewapply(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewapplyAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionChangeNewapplyAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionChangeNewapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
