package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundApply 商旅机票分销-退票申请
// alitrip.btrip.flight.distribution.refund.apply
//
// 商旅机票分销-退票申请
func AlitripBtripFlightDistributionRefundApply(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundApplyAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
