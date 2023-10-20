package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinateprocessing 慧飞商家协同单处理完成接口
// alitrip.agent.coordinate.processing
//
// 慧飞商家协同单处理完成接口
func Alitripagentcoordinateprocessing(clt *core.SDKClient, req *flight.AlitripagentcoordinateprocessingAPIRequest, session string) (*flight.AlitripagentcoordinateprocessingAPIResponse, error) {
	var resp flight.AlitripagentcoordinateprocessingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
