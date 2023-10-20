package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinateprocess 慧飞商家协同单处理完成接口
// alitrip.agent.coordinate.process
//
// 慧飞商家协同单处理完成接口
func Alitripagentcoordinateprocess(clt *core.SDKClient, req *flight.AlitripagentcoordinateprocessAPIRequest, session string) (*flight.AlitripagentcoordinateprocessAPIResponse, error) {
	var resp flight.AlitripagentcoordinateprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
