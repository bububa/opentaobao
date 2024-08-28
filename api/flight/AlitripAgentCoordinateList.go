package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateList 慧飞商家协同单列表查询接口
// alitrip.agent.coordinate.list
//
// 慧飞商家协同单列表查询接口
func AlitripAgentCoordinateList(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateListAPIRequest, resp *flight.AlitripAgentCoordinateListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
