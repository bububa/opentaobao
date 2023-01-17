package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateProcess 慧飞商家协同单处理完成接口
// alitrip.agent.coordinate.process
//
// 慧飞商家协同单处理完成接口
func AlitripAgentCoordinateProcess(clt *core.SDKClient, req *flight.AlitripAgentCoordinateProcessAPIRequest, session string) (*flight.AlitripAgentCoordinateProcessAPIResponse, error) {
	var resp flight.AlitripAgentCoordinateProcessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
