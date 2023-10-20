package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellmodifyapprove 销售改签确认
// alitrip.agent.flight.sell.modify.approve
//
// 销售改签确认
func Alitripagentflightsellmodifyapprove(clt *core.SDKClient, req *flight.AlitripagentflightsellmodifyapproveAPIRequest, session string) (*flight.AlitripagentflightsellmodifyapproveAPIResponse, error) {
	var resp flight.AlitripagentflightsellmodifyapproveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
