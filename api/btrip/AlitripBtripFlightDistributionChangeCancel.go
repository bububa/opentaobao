package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeCancel 商旅机票改签取消
// alitrip.btrip.flight.distribution.change.cancel
//
// 商旅机票改签取消
func AlitripBtripFlightDistributionChangeCancel(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeCancelAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
