package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionModifyFlightsearch 改签航班列表
// alitrip.btrip.flight.distribution.modify.flightsearch
//
// 商旅分销改签航班列表
func AlitripBtripFlightDistributionModifyFlightsearch(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionModifyFlightsearchAPIRequest, resp *btrip.AlitripBtripFlightDistributionModifyFlightsearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
