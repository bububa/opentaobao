package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundNewapply 商旅机票分销-退票申请
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
func AlitripBtripFlightDistributionRefundNewapply(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundNewapplyAPIRequest, session string) (*btrip.AlitripBtripFlightDistributionRefundNewapplyAPIResponse, error) {
	var resp btrip.AlitripBtripFlightDistributionRefundNewapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
