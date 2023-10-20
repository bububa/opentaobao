package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundDetail 销售退票单详情
// alitrip.agent.flight.sell.refund.detail
//
// 销售退票单详情
func AlitripAgentFlightSellRefundDetail(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundDetailAPIRequest, resp *flight.AlitripAgentFlightSellRefundDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
