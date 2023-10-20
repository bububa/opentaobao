package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellticketinglist 销售出票列表
// alitrip.agent.flight.sell.ticketing.list
//
// 销售出票列表
func Alitripagentflightsellticketinglist(clt *core.SDKClient, req *flight.AlitripagentflightsellticketinglistAPIRequest, session string) (*flight.AlitripagentflightsellticketinglistAPIResponse, error) {
	var resp flight.AlitripagentflightsellticketinglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
