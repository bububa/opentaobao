package ieagency

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ieagency"
)

// TaobaoAlitripIeAgentRefundRefundmoney 确认退款
// taobao.alitrip.ie.agent.refund.refundmoney
//
// 卖家进行退款操作
func TaobaoAlitripIeAgentRefundRefundmoney(clt *core.SDKClient, req *ieagency.TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest, resp *ieagency.TaobaoAlitripIeAgentRefundRefundmoneyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
