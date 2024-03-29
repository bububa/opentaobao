package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentOrderGet 【国际机票】查询订单详情
// taobao.alitrip.ie.agent.order.get
//
// 根据订单ID查询订单详情
func TaobaoAlitripIeAgentOrderGet(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentOrderGetAPIRequest, resp *ieagency.TaobaoAlitripIeAgentOrderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
