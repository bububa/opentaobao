package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateDetail 商家协同单查询详情
// alitrip.agent.coordinate.detail
//
// 商家协同单查询详情
func AlitripAgentCoordinateDetail(clt *core.SDKClient, req *flight.AlitripAgentCoordinateDetailAPIRequest, session string) (*flight.AlitripAgentCoordinateDetailAPIResponse, error) {
	var resp flight.AlitripAgentCoordinateDetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
