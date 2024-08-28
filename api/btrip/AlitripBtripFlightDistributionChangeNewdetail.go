package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionChangeNewdetail 商旅机票改签详情接口
// alitrip.btrip.flight.distribution.change.newdetail
//
// 商旅机票改签详情接口
func AlitripBtripFlightDistributionChangeNewdetail(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionChangeNewdetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionChangeNewdetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
