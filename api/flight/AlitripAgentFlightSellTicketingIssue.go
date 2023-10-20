package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellticketingissue 销售出票
// alitrip.agent.flight.sell.ticketing.issue
//
// 销售出票
func Alitripagentflightsellticketingissue(clt *core.SDKClient, req *flight.AlitripagentflightsellticketingissueAPIRequest, session string) (*flight.AlitripagentflightsellticketingissueAPIResponse, error) {
	var resp flight.AlitripagentflightsellticketingissueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
