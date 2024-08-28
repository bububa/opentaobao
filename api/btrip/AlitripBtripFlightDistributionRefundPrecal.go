package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundPrecal 商旅机票分销-退票费预计算
// alitrip.btrip.flight.distribution.refund.precal
//
// 商旅机票分销-退票费预计算
func AlitripBtripFlightDistributionRefundPrecal(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundPrecalAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundPrecalAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
