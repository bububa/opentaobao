package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateDetail 商家协同单查询详情
// alitrip.agent.coordinate.detail
//
// 商家协同单查询详情
func AlitripAgentCoordinateDetail(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateDetailAPIRequest, resp *flight.AlitripAgentCoordinateDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
