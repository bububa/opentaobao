package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateHandle 慧飞商家协同单接手接口
// alitrip.agent.coordinate.handle
//
// 慧飞商家协同单接手接口
func AlitripAgentCoordinateHandle(clt *core.SDKClient, req *flight.AlitripAgentCoordinateHandleAPIRequest, resp *flight.AlitripAgentCoordinateHandleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
