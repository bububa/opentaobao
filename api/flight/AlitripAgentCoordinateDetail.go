package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentcoordinatedetail 商家协同单查询详情
// alitrip.agent.coordinate.detail
//
// 商家协同单查询详情
func Alitripagentcoordinatedetail(clt *core.SDKClient, req *flight.AlitripagentcoordinatedetailAPIRequest, session string) (*flight.AlitripagentcoordinatedetailAPIResponse, error) {
	var resp flight.AlitripagentcoordinatedetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
