package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellmodifyrefuse 销售改签拒绝
// alitrip.agent.flight.sell.modify.refuse
//
// 销售改签拒绝
func Alitripagentflightsellmodifyrefuse(clt *core.SDKClient, req *flight.AlitripagentflightsellmodifyrefuseAPIRequest, session string) (*flight.AlitripagentflightsellmodifyrefuseAPIResponse, error) {
	var resp flight.AlitripagentflightsellmodifyrefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
