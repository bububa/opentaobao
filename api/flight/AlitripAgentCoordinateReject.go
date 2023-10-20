package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinatereject 慧飞商家协同单拒绝接口
// alitrip.agent.coordinate.reject
//
// 慧飞商家协同单拒绝接口
func Alitripagentcoordinatereject(clt *core.SDKClient, req *flight.AlitripagentcoordinaterejectAPIRequest, session string) (*flight.AlitripagentcoordinaterejectAPIResponse, error) {
	var resp flight.AlitripagentcoordinaterejectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
