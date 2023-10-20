package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// AlitripTripvpAgentOrderSearch 【国际机票】查询辅营订单列表
// alitrip.tripvp.agent.order.search
//
// 【国际机票】查询辅营订单列表
func AlitripTripvpAgentOrderSearch(clt *core.SDKClient, req *ieagency.AlitripTripvpAgentOrderSearchAPIRequest, resp *ieagency.AlitripTripvpAgentOrderSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
