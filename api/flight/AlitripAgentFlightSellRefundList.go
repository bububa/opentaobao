package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundList 销售退票单列表
// alitrip.agent.flight.sell.refund.list
//
// 销售退票单列表
func AlitripAgentFlightSellRefundList(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundListAPIRequest, resp *flight.AlitripAgentFlightSellRefundListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
