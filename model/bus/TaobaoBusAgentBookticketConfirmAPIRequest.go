package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentBookticketConfirmAPIRequest 汽车票代理商接口—确认出票是否成功 API请求
// taobao.bus.agent.bookticket.confirm
//
// 代理商通过该接口通知汽车票系统订单出票结果。
type TaobaoBusAgentBookticketConfirmAPIRequest struct {
	model.Params
	// 请求对象
	_paramAgentConfirmBookRQ *AgentConfirmBookRq
}

// NewTaobaoBusAgentBookticketConfirmRequest 初始化TaobaoBusAgentBookticketConfirmAPIRequest对象
func NewTaobaoBusAgentBookticketConfirmRequest() *TaobaoBusAgentBookticketConfirmAPIRequest {
	return &TaobaoBusAgentBookticketConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentBookticketConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.bookticket.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentBookticketConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamAgentConfirmBookRQ is ParamAgentConfirmBookRQ Setter
// 请求对象
func (r *TaobaoBusAgentBookticketConfirmAPIRequest) SetParamAgentConfirmBookRQ(_paramAgentConfirmBookRQ *AgentConfirmBookRq) error {
	r._paramAgentConfirmBookRQ = _paramAgentConfirmBookRQ
	r.Set("param_agent_confirm_book_r_q", _paramAgentConfirmBookRQ)
	return nil
}

// GetParamAgentConfirmBookRQ ParamAgentConfirmBookRQ Getter
func (r TaobaoBusAgentBookticketConfirmAPIRequest) GetParamAgentConfirmBookRQ() *AgentConfirmBookRq {
	return r._paramAgentConfirmBookRQ
}
