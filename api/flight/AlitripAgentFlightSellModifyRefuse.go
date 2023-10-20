package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyRefuse 销售改签拒绝
// alitrip.agent.flight.sell.modify.refuse
//
// 销售改签拒绝
func AlitripAgentFlightSellModifyRefuse(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyRefuseAPIRequest, resp *flight.AlitripAgentFlightSellModifyRefuseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
