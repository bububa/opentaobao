package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateReject 慧飞商家协同单拒绝接口
// alitrip.agent.coordinate.reject
//
// 慧飞商家协同单拒绝接口
func AlitripAgentCoordinateReject(clt *core.SDKClient, req *flight.AlitripAgentCoordinateRejectAPIRequest, resp *flight.AlitripAgentCoordinateRejectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
