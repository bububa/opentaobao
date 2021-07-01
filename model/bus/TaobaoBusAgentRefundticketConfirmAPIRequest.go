package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusAgentRefundticketConfirmAPIRequest
商家top回调退款明细 API请求
taobao.bus.agent.refundticket.confirm

商家通过top回调告知平台退款明细 */
type TaobaoBusAgentRefundticketConfirmAPIRequest struct {
	model.Params
	// 退款入参
	_paramAgentConfirmRefundRQ *AgentConfirmRefundRq
}

// New
