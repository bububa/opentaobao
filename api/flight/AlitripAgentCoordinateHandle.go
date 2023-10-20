package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinatehandle 慧飞商家协同单接手接口
// alitrip.agent.coordinate.handle
//
// 慧飞商家协同单接手接口
func Alitripagentcoordinatehandle(clt *core.SDKClient, req *flight.AlitripagentcoordinatehandleAPIRequest, session string) (*flight.AlitripagentcoordinatehandleAPIResponse, error) {
	var resp flight.AlitripagentcoordinatehandleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
