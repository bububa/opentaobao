package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightIntentionList 意向单列表
// alitrip.agent.flight.intention.list
//
// 意向单列表
func AlitripAgentFlightIntentionList(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentFlightIntentionListAPIRequest, resp *flight.AlitripAgentFlightIntentionListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
