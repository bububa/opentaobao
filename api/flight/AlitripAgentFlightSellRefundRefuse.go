package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellrefundrefuse 销售退票拒绝
// alitrip.agent.flight.sell.refund.refuse
//
// 销售退票拒绝
func Alitripagentflightsellrefundrefuse(clt *core.SDKClient, req *flight.AlitripagentflightsellrefundrefuseAPIRequest, session string) (*flight.AlitripagentflightsellrefundrefuseAPIResponse, error) {
	var resp flight.AlitripagentflightsellrefundrefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
