package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundRefuse 销售退票拒绝
// alitrip.agent.flight.sell.refund.refuse
//
// 销售退票拒绝
func AlitripAgentFlightSellRefundRefuse(clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundRefuseAPIRequest, resp *flight.AlitripAgentFlightSellRefundRefuseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
