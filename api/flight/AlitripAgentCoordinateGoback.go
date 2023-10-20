package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinategoback 协同单驳回
// alitrip.agent.coordinate.goback
//
// 协同单驳回
func Alitripagentcoordinategoback(clt *core.SDKClient, req *flight.AlitripagentcoordinategobackAPIRequest, session string) (*flight.AlitripagentcoordinategobackAPIResponse, error) {
	var resp flight.AlitripagentcoordinategobackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
