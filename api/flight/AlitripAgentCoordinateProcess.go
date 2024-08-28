package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateProcess 慧飞商家协同单处理完成接口
// alitrip.agent.coordinate.process
//
// 慧飞商家协同单处理完成接口
func AlitripAgentCoordinateProcess(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateProcessAPIRequest, resp *flight.AlitripAgentCoordinateProcessAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
