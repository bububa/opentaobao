package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateReject 慧飞商家协同单拒绝接口
// alitrip.agent.coordinate.reject
//
// 慧飞商家协同单拒绝接口
func AlitripAgentCoordinateReject(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateRejectAPIRequest, resp *flight.AlitripAgentCoordinateRejectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
