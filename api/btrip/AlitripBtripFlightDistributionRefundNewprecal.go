package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundNewprecal 商旅机票分销-退票费预计算
// alitrip.btrip.flight.distribution.refund.newprecal
//
// 商旅机票分销-退票费预计算
func AlitripBtripFlightDistributionRefundNewprecal(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundNewprecalAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundNewprecalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
