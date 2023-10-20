package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundList 销售退票单列表
// alitrip.agent.flight.sell.refund.list
//
// 销售退票单列表
func AlitripAgentFlightSellRefundList(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundListAPIRequest, resp *flight.AlitripAgentFlightSellRefundListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
