package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeApply 商旅机票改签申请
// alitrip.btrip.flight.distribution.change.apply
//
// 商旅机票改签申请
func AlitripBtripFlightDistributionChangeApply(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeApplyAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeApplyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
