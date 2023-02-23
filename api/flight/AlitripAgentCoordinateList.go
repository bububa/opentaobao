package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentCoordinateList 慧飞商家协同单列表查询接口
// alitrip.agent.coordinate.list
//
// 慧飞商家协同单列表查询接口
func AlitripAgentCoordinateList(clt *core.SDKClient, req *flight.AlitripAgentCoordinateListAPIRequest, session string) (*flight.AlitripAgentCoordinateListAPIResponse, error) {
	var resp flight.AlitripAgentCoordinateListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
