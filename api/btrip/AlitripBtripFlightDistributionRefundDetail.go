package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionRefundDetail 商旅机票退票详情接口
// alitrip.btrip.flight.distribution.refund.detail
//
// 商旅机票分销退票详情
func AlitripBtripFlightDistributionRefundDetail(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionRefundDetailAPIRequest, resp *btrip.AlitripBtripFlightDistributionRefundDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
