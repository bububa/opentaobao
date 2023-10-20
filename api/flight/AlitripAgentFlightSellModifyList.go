package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellmodifylist 销售改签单列表
// alitrip.agent.flight.sell.modify.list
//
// 销售改签单列表
func Alitripagentflightsellmodifylist(clt *core.SDKClient, req *flight.AlitripagentflightsellmodifylistAPIRequest, session string) (*flight.AlitripagentflightsellmodifylistAPIResponse, error) {
	var resp flight.AlitripagentflightsellmodifylistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
