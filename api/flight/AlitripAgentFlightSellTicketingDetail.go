package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellticketingdetail 销售出票详情
// alitrip.agent.flight.sell.ticketing.detail
//
// 销售出票详情
func Alitripagentflightsellticketingdetail(clt *core.SDKClient, req *flight.AlitripagentflightsellticketingdetailAPIRequest, session string) (*flight.AlitripagentflightsellticketingdetailAPIResponse, error) {
	var resp flight.AlitripagentflightsellticketingdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
