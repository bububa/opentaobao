package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitriptripvpagentorderissue 廉航辅营正向订单出货接口
// alitrip.tripvp.agent.order.issue
//
// 廉航辅营正向订单出货接口
func Alitriptripvpagentorderissue(clt *core.SDKClient, req *flight.AlitriptripvpagentorderissueAPIRequest, session string) (*flight.AlitriptripvpagentorderissueAPIResponse, error) {
	var resp flight.AlitriptripvpagentorderissueAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
