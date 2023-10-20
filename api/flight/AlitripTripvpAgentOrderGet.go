package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitriptripvpagentorderget 廉航辅营正向订单查询详情接口
// alitrip.tripvp.agent.order.get
//
// 【国际机票】查询辅营订单详情
func Alitriptripvpagentorderget(clt *core.SDKClient, req *flight.AlitriptripvpagentordergetAPIRequest, session string) (*flight.AlitriptripvpagentordergetAPIResponse, error) {
	var resp flight.AlitriptripvpagentordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
