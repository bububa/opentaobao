package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellRefundRefuse 销售退票拒绝
// alitrip.agent.flight.sell.refund.refuse
//
// 销售退票拒绝
func AlitripAgentFlightSellRefundRefuse(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellRefundRefuseAPIRequest, resp *flight.AlitripAgentFlightSellRefundRefuseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
