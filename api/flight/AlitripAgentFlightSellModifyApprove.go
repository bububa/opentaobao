package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightSellModifyApprove 销售改签确认
// alitrip.agent.flight.sell.modify.approve
//
// 销售改签确认
func AlitripAgentFlightSellModifyApprove(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightSellModifyApproveAPIRequest, resp *flight.AlitripAgentFlightSellModifyApproveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
