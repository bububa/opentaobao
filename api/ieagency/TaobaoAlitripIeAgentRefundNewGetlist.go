package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundNewGetlist 新查询退票申请单列表
// taobao.alitrip.ie.agent.refund.new.getlist
//
// 查询商家国际机票退票申请单列表
func TaobaoAlitripIeAgentRefundNewGetlist(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundNewGetlistAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundNewGetlistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
