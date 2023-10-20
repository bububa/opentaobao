package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellrefundapprove 销售退票确认
// alitrip.agent.flight.sell.refund.approve
//
// 销售退票确认
func Alitripagentflightsellrefundapprove(clt *core.SDKClient, req *flight.AlitripagentflightsellrefundapproveAPIRequest, session string) (*flight.AlitripagentflightsellrefundapproveAPIResponse, error) {
	var resp flight.AlitripagentflightsellrefundapproveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
