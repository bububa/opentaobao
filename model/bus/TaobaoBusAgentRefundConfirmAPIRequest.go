package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusagentrefundconfirmAPIRequest 汽车票退票和退款二合一接口 API请求
// taobao.bus.agent.refund.confirm
//
// 1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
type TaobaobusagentrefundconfirmAPIRequest struct {
	model.Params
	// 入参
	_paramAgentConfirmReturnAndRefundRQ *AgentConfirmReturnAndRefundRq
}

// NewTaobaobusagentrefundconfirmRequest 初始化TaobaobusagentrefundconfirmAPIRequest对象
func NewTaobaobusagentrefundconfirmRequest() *TaobaobusagentrefundconfirmAPIRequest {
	return &TaobaobusagentrefundconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusagentrefundconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.refund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusagentrefundconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusagentrefundconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentConfirmReturnAndRefundRQ is ParamAgentConfirmReturnAndRefundRQ Setter
// 入参
func (r *TaobaobusagentrefundconfirmAPIRequest) SetParamAgentConfirmReturnAndRefundRQ(_paramAgentConfirmReturnAndRefundRQ *AgentConfirmReturnAndRefundRq) error {
	r._paramAgentConfirmReturnAndRefundRQ = _paramAgentConfirmReturnAndRefundRQ
	r.Set("param_agent_confirm_return_and_refund_r_q", _paramAgentConfirmReturnAndRefundRQ)
	return nil
}

// GetParamAgentConfirmReturnAndRefundRQ ParamAgentConfirmReturnAndRefundRQ Getter
func (r TaobaobusagentrefundconfirmAPIRequest) GetParamAgentConfirmReturnAndRefundRQ() *AgentConfirmReturnAndRefundRq {
	return r._paramAgentConfirmReturnAndRefundRQ
}
