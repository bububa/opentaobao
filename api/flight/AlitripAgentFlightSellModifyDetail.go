package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellmodifydetail 销售改签详情
// alitrip.agent.flight.sell.modify.detail
//
// 销售改签详情
func Alitripagentflightsellmodifydetail(clt *core.SDKClient, req *flight.AlitripagentflightsellmodifydetailAPIRequest, session string) (*flight.AlitripagentflightsellmodifydetailAPIResponse, error) {
	var resp flight.AlitripagentflightsellmodifydetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
