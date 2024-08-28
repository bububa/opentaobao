package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundApply 商旅机票分销-退票申请
// alitrip.btrip.flight.distribution.refund.apply
//
// 商旅机票分销-退票申请
func AlitripBtripFlightDistributionRefundApply(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundApplyAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
