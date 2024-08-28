package btrip

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundDetail 商旅机票退票详情接口
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
func AlitripBtripFlightDistributionRefundDetail(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundDetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
