package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobusagentbookticketconfirmAPIRequest 汽车票代理商接口—确认出票是否成功 API请求
// taobao.bus.agent.bookticket.confirm
//
// 代理商通过该接口通知汽车票系统订单出票结果。
type TaobaobusagentbookticketconfirmAPIRequest struct {
	model.Params
	// 请求对象
	_paramAgentConfirmBookRQ *AgentConfirmBookRq
}

// NewTaobaobusagentbookticketconfirmRequest 初始化TaobaobusagentbookticketconfirmAPIRequest对象
func NewTaobaobusagentbookticketconfirmRequest() *TaobaobusagentbookticketconfirmAPIRequest {
	return &TaobaobusagentbookticketconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobusagentbookticketconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.bookticket.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobusagentbookticketconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobusagentbookticketconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentConfirmBookRQ is ParamAgentConfirmBookRQ Setter
// 请求对象
func (r *TaobaobusagentbookticketconfirmAPIRequest) SetParamAgentConfirmBookRQ(_paramAgentConfirmBookRQ *AgentConfirmBookRq) error {
	r._paramAgentConfirmBookRQ = _paramAgentConfirmBookRQ
	r.Set("param_agent_confirm_book_r_q", _paramAgentConfirmBookRQ)
	return nil
}

// GetParamAgentConfirmBookRQ ParamAgentConfirmBookRQ Getter
func (r TaobaobusagentbookticketconfirmAPIRequest) GetParamAgentConfirmBookRQ() *AgentConfirmBookRq {
	return r._paramAgentConfirmBookRQ
}
