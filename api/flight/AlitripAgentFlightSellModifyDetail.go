package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyDetail 销售改签详情
// alitrip.agent.flight.sell.modify.detail
//
// 销售改签详情
func AlitripAgentFlightSellModifyDetail(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyDetailAPIRequest, resp *flight.AlitripAgentFlightSellModifyDetailAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
