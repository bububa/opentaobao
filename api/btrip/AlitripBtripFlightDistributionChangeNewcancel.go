package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewcancel 商旅机票改签取消
// alitrip.btrip.flight.distribution.change.newcancel
//
// 商旅机票改签取消
func AlitripBtripFlightDistributionChangeNewcancel(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewcancelAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeNewcancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
