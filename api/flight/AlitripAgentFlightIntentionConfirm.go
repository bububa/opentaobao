package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightIntentionConfirm 意向单确认
// alitrip.agent.flight.intention.confirm
//
// 意向单确认
func AlitripAgentFlightIntentionConfirm(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightIntentionConfirmAPIRequest, resp *flight.AlitripAgentFlightIntentionConfirmAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
