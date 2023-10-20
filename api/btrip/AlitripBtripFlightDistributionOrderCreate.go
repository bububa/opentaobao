package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripFlightDistributionOrderCreate 商旅机票分销-创建订单
// alitrip.btrip.flight.distribution.order.create
//
// 商旅机票分销创建订单接口
func AlitripBtripFlightDistributionOrderCreate(clt *core.SDKClient, req *btrip.AlitripBtripFlightDistributionOrderCreateAPIRequest, resp *btrip.AlitripBtripFlightDistributionOrderCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
