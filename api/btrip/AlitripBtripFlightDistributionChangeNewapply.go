package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewapply 商旅机票改签申请V2
// alitrip.btrip.flight.distribution.change.newapply
//
// 商旅机票改签申请
func AlitripBtripFlightDistributionChangeNewapply(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewapplyAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeNewapplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
