package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellrefundlist 销售退票单列表
// alitrip.agent.flight.sell.refund.list
//
// 销售退票单列表
func Alitripagentflightsellrefundlist(clt *core.SDKClient, req *flight.AlitripagentflightsellrefundlistAPIRequest, session string) (*flight.AlitripagentflightsellrefundlistAPIResponse, error) {
	var resp flight.AlitripagentflightsellrefundlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
