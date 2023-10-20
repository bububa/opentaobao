package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightsellmodifybackfill 销售改签回填
// alitrip.agent.flight.sell.modify.backfill
//
// 销售改签回填
func Alitripagentflightsellmodifybackfill(clt *core.SDKClient, req *flight.AlitripagentflightsellmodifybackfillAPIRequest, session string) (*flight.AlitripagentflightsellmodifybackfillAPIResponse, error) {
	var resp flight.AlitripagentflightsellmodifybackfillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
