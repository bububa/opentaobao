package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyBackfill 销售改签回填
// alitrip.agent.flight.sell.modify.backfill
//
// 销售改签回填
func AlitripAgentFlightSellModifyBackfill(clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyBackfillAPIRequest, resp *flight.AlitripAgentFlightSellModifyBackfillAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
