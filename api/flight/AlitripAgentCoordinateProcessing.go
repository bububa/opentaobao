package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateProcessing 慧飞商家协同单处理完成接口
// alitrip.agent.coordinate.processing
//
// 慧飞商家协同单处理完成接口
func AlitripAgentCoordinateProcessing(clt *core.SDKClient, req *flight.AlitripAgentCoordinateProcessingAPIRequest, resp *flight.AlitripAgentCoordinateProcessingAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
