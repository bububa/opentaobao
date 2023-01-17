package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateGoback 协同单驳回
// alitrip.agent.coordinate.goback
//
// 协同单驳回
func AlitripAgentCoordinateGoback(clt *core.SDKClient, req *flight.AlitripAgentCoordinateGobackAPIRequest, session string) (*flight.AlitripAgentCoordinateGobackAPIResponse, error) {
	var resp flight.AlitripAgentCoordinateGobackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
