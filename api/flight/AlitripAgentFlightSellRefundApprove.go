package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundApprove 销售退票确认
// alitrip.agent.flight.sell.refund.approve
//
// 销售退票确认
func AlitripAgentFlightSellRefundApprove(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundApproveAPIRequest, resp *flight.AlitripAgentFlightSellRefundApproveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
