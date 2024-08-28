package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundNewapply 商旅机票分销-退票申请
// alitrip.btrip.flight.distribution.refund.newapply
//
// 商旅机票分销-退票申请
func AlitripBtripFlightDistributionRefundNewapply(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundNewapplyAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundNewapplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
