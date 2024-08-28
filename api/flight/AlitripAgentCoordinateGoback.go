package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateGoback 协同单驳回
// alitrip.agent.coordinate.goback
//
// 协同单驳回
func AlitripAgentCoordinateGoback(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateGobackAPIRequest, resp *flight.AlitripAgentCoordinateGobackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
