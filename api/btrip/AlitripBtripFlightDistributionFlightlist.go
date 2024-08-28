package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionFlightlist 商旅机票航班列表接口
// alitrip.btrip.flight.distribution.flightlist
//
// 商旅机票航班列表接口，用于分销询价
func AlitripBtripFlightDistributionFlightlist(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionFlightlistAPIRequest, resp *btrip.AlitripBtripFlightDistributionFlightlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
