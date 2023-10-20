package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusagentrefundticketconfirmAPIRequest 商家top回调退款明细 API请求
// taobao.bus.agent.refundticket.confirm
//
// 商家通过top回调告知平台退款明细
type TaobaobusagentrefundticketconfirmAPIRequest struct {
	model.Params
	// 退款入参
	_paramAgentConfirmRefundRQ *AgentConfirmRefundRq
}

// NewTaobaobusagentrefundticketconfirmRequest 初始化TaobaobusagentrefundticketconfirmAPIRequest对象
func NewTaobaobusagentrefundticketconfirmRequest() *TaobaobusagentrefundticketconfirmAPIRequest {
	return &TaobaobusagentrefundticketconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusagentrefundticketconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.refundticket.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusagentrefundticketconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusagentrefundticketconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentConfirmRefundRQ is ParamAgentConfirmRefundRQ Setter
// 退款入参
func (r *TaobaobusagentrefundticketconfirmAPIRequest) SetParamAgentConfirmRefundRQ(_paramAgentConfirmRefundRQ *AgentConfirmRefundRq) error {
	r._paramAgentConfirmRefundRQ = _paramAgentConfirmRefundRQ
	r.Set("param_agent_confirm_refund_r_q", _paramAgentConfirmRefundRQ)
	return nil
}

// GetParamAgentConfirmRefundRQ ParamAgentConfirmRefundRQ Getter
func (r TaobaobusagentrefundticketconfirmAPIRequest) GetParamAgentConfirmRefundRQ() *AgentConfirmRefundRq {
	return r._paramAgentConfirmRefundRQ
}
