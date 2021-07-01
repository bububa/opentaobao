package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest
确认退款 API请求
taobao.alitrip.ie.agent.refund.refundmoney

卖家进行退款操作 */
type TaobaoAlitripIeAgentRefundRefundmoneyAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
}

// New
