package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripTripvpAgentOrderGet 廉航辅营正向订单查询详情接口
// alitrip.tripvp.agent.order.get
//
// 【国际机票】查询辅营订单详情
func AlitripTripvpAgentOrderGet(clt *core.SDKClient, req *flight.AlitripTripvpAgentOrderGetAPIRequest, resp *flight.AlitripTripvpAgentOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
