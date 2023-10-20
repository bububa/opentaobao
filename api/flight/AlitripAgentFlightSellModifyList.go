package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyList 销售改签单列表
// alitrip.agent.flight.sell.modify.list
//
// 销售改签单列表
func AlitripAgentFlightSellModifyList(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyListAPIRequest, resp *flight.AlitripAgentFlightSellModifyListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
