package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyDetail 销售改签详情
// alitrip.agent.flight.sell.modify.detail
//
// 销售改签详情
func AlitripAgentFlightSellModifyDetail(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyDetailAPIRequest, resp *flight.AlitripAgentFlightSellModifyDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
