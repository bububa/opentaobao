package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundAgreeAPIRequest
同意退票 API请求
taobao.alitrip.ie.agent.refund.agree

卖家同意买家退票申请 */
type TaobaoAlitripIeAgentRefundAgreeAPIRequest struct {
	model.Params
	// 退款金额
	_refundMoney int64
	// 申请单id
	_applyId int64
	// 订单id
	_orderId int64
	// 回复信息
	_agentAnswer string
	// 代理商id
	_agentId int64
}

// New
