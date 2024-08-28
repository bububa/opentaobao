package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionNewflightlist 商旅机票航班列表接口，用于分销询价V2
// alitrip.btrip.flight.distribution.newflightlist
//
// 商旅机票航班列表接口，用于分销询价V2
func AlitripBtripFlightDistributionNewflightlist(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionNewflightlistAPIRequest, resp *btrip.AlitripBtripFlightDistributionNewflightlistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
