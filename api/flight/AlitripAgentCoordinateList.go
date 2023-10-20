package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinatelist 慧飞商家协同单列表查询接口
// alitrip.agent.coordinate.list
//
// 慧飞商家协同单列表查询接口
func Alitripagentcoordinatelist(clt *core.SDKClient, req *flight.AlitripagentcoordinatelistAPIRequest, session string) (*flight.AlitripagentcoordinatelistAPIResponse, error) {
	var resp flight.AlitripagentcoordinatelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
