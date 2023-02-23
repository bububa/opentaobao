package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateProcessing 慧飞商家协同单处理完成接口
// alitrip.agent.coordinate.processing
//
// 慧飞商家协同单处理完成接口
func AlitripAgentCoordinateProcessing(clt *core.SDKClient, req *flight.AlitripAgentCoordinateProcessingAPIRequest, session string) (*flight.AlitripAgentCoordinateProcessingAPIResponse, error) {
	var resp flight.AlitripAgentCoordinateProcessingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
