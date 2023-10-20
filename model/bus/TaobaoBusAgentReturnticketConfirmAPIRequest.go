package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusagentreturnticketconfirmAPIRequest 商家回调退票 API请求
// taobao.bus.agent.returnticket.confirm
//
// 商家通过TOP接口调用来回传退票状态
type TaobaobusagentreturnticketconfirmAPIRequest struct {
	model.Params
	// 退票入参
	_paramAgentConfirmReturnRQ *AgentConfirmReturnRq
}

// NewTaobaobusagentreturnticketconfirmRequest 初始化TaobaobusagentreturnticketconfirmAPIRequest对象
func NewTaobaobusagentreturnticketconfirmRequest() *TaobaobusagentreturnticketconfirmAPIRequest {
	return &TaobaobusagentreturnticketconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusagentreturnticketconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.returnticket.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusagentreturnticketconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusagentreturnticketconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentConfirmReturnRQ is ParamAgentConfirmReturnRQ Setter
// 退票入参
func (r *TaobaobusagentreturnticketconfirmAPIRequest) SetParamAgentConfirmReturnRQ(_paramAgentConfirmReturnRQ *AgentConfirmReturnRq) error {
	r._paramAgentConfirmReturnRQ = _paramAgentConfirmReturnRQ
	r.Set("param_agent_confirm_return_r_q", _paramAgentConfirmReturnRQ)
	return nil
}

// GetParamAgentConfirmReturnRQ ParamAgentConfirmReturnRQ Getter
func (r TaobaobusagentreturnticketconfirmAPIRequest) GetParamAgentConfirmReturnRQ() *AgentConfirmReturnRq {
	return r._paramAgentConfirmReturnRQ
}
