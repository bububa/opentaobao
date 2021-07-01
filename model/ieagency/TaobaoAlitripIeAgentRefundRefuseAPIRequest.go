package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundRefuseAPIRequest
拒绝退票申请 API请求
taobao.alitrip.ie.agent.refund.refuse

卖家拒绝退票退票申请 */
type TaobaoAlitripIeAgentRefundRefuseAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 订单id
	_orderId int64
	// 代理商回复
	_agentAnswer string
	// 代理商id
	_agentId int64
}

// New
