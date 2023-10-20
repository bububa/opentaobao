package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundRefuse 拒绝退票申请
// taobao.alitrip.ie.agent.refund.refuse
//
// 卖家拒绝退票退票申请
func TaobaoAlitripIeAgentRefundRefuse(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundRefuseAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundRefuseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
