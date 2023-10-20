package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentOrderSearch 【国际机票】订单列表查询
// taobao.alitrip.ie.agent.order.search
//
// 根据指定条件查询国际机票订单列表
func TaobaoAlitripIeAgentOrderSearch(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentOrderSearchAPIRequest, resp *ieagency.TaobaoAlitripIeAgentOrderSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
