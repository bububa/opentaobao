package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionModifyNewflightsearch 改签航班列表V2
// alitrip.btrip.flight.distribution.modify.newflightsearch
//
// 改签航班列表V2
func AlitripBtripFlightDistributionModifyNewflightsearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionModifyNewflightsearchAPIRequest, resp *btrip.AlitripBtripFlightDistributionModifyNewflightsearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
