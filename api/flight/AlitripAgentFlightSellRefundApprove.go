package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundApprove 销售退票确认
// alitrip.agent.flight.sell.refund.approve
//
// 销售退票确认
func AlitripAgentFlightSellRefundApprove(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundApproveAPIRequest, resp *flight.AlitripAgentFlightSellRefundApproveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
