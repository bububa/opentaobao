package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateProcessing 慧飞商家协同单处理完成接口
// alitrip.agent.coordinate.processing
//
// 慧飞商家协同单处理完成接口
func AlitripAgentCoordinateProcessing(ctx context.Context, clt *core.SDKClient, req *flight.AlitripAgentCoordinateProcessingAPIRequest, resp *flight.AlitripAgentCoordinateProcessingAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
