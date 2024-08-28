package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewquery 改签航班询价V2
// alitrip.btrip.flight.distribution.change.newquery
//
// 商旅机票改签航班询价
func AlitripBtripFlightDistributionChangeNewquery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewqueryAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeNewqueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
