package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellrefunddetail 销售退票单详情
// alitrip.agent.flight.sell.refund.detail
//
// 销售退票单详情
func Alitripagentflightsellrefunddetail(clt *core.SDKClient, req *flight.AlitripagentflightsellrefunddetailAPIRequest, session string) (*flight.AlitripagentflightsellrefunddetailAPIResponse, error) {
	var resp flight.AlitripagentflightsellrefunddetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
