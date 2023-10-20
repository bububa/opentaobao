package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyBackfill 销售改签回填
// alitrip.agent.flight.sell.modify.backfill
//
// 销售改签回填
func AlitripAgentFlightSellModifyBackfill(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyBackfillAPIRequest, session string) (*flight.AlitripAgentFlightSellModifyBackfillAPIResponse, error) {
	var resp flight.AlitripAgentFlightSellModifyBackfillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
