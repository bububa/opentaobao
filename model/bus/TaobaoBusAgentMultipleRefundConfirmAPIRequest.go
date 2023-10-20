package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusagentmultiplerefundconfirmAPIRequest 综合交通多次退款接口 API请求
// taobao.bus.agent.multiple.refund.confirm
//
// 此接口支持多次按照单客进行多次退款操作，只进行退款操作。
type TaobaobusagentmultiplerefundconfirmAPIRequest struct {
	model.Params
	// 入参
	_paramAgentMultipleRefundRQ *AgentMultipleRefundRq
}

// NewTaobaobusagentmultiplerefundconfirmRequest 初始化TaobaobusagentmultiplerefundconfirmAPIRequest对象
func NewTaobaobusagentmultiplerefundconfirmRequest() *TaobaobusagentmultiplerefundconfirmAPIRequest {
	return &TaobaobusagentmultiplerefundconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusagentmultiplerefundconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.multiple.refund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusagentmultiplerefundconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusagentmultiplerefundconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentMultipleRefundRQ is ParamAgentMultipleRefundRQ Setter
// 入参
func (r *TaobaobusagentmultiplerefundconfirmAPIRequest) SetParamAgentMultipleRefundRQ(_paramAgentMultipleRefundRQ *AgentMultipleRefundRq) error {
	r._paramAgentMultipleRefundRQ = _paramAgentMultipleRefundRQ
	r.Set("param_agent_multiple_refund_r_q", _paramAgentMultipleRefundRQ)
	return nil
}

// GetParamAgentMultipleRefundRQ ParamAgentMultipleRefundRQ Getter
func (r TaobaobusagentmultiplerefundconfirmAPIRequest) GetParamAgentMultipleRefundRQ() *AgentMultipleRefundRq {
	return r._paramAgentMultipleRefundRQ
}
