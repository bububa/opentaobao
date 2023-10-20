package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderCancel 商旅机票分销-取消订单
// alitrip.btrip.flight.distribution.order.cancel
//
// 商旅机票分销取消订单
func AlitripBtripFlightDistributionOrderCancel(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderCancelAPIRequest, resp *btrip.AlitripBtripFlightDistributionOrderCancelAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
