package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentRefundticketConfirmAPIRequest 商家top回调退款明细 API请求
// taobao.bus.agent.refundticket.confirm
//
// 商家通过top回调告知平台退款明细
type TaobaoBusAgentRefundticketConfirmAPIRequest struct {
	model.Params
	// 退款入参
	_paramAgentConfirmRefundRQ *AgentConfirmRefundRq
}

// NewTaobaoBusAgentRefundticketConfirmRequest 初始化TaobaoBusAgentRefundticketConfirmAPIRequest对象
func NewTaobaoBusAgentRefundticketConfirmRequest() *TaobaoBusAgentRefundticketConfirmAPIRequest {
	return &TaobaoBusAgentRefundticketConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusAgentRefundticketConfirmAPIRequest) Reset() {
	r._paramAgentConfirmRefundRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentRefundticketConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.refundticket.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentRefundticketConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusAgentRefundticketConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentConfirmRefundRQ is ParamAgentConfirmRefundRQ Setter
// 退款入参
func (r *TaobaoBusAgentRefundticketConfirmAPIRequest) SetParamAgentConfirmRefundRQ(_paramAgentConfirmRefundRQ *AgentConfirmRefundRq) error {
	r._paramAgentConfirmRefundRQ = _paramAgentConfirmRefundRQ
	r.Set("param_agent_confirm_refund_r_q", _paramAgentConfirmRefundRQ)
	return nil
}

// GetParamAgentConfirmRefundRQ ParamAgentConfirmRefundRQ Getter
func (r TaobaoBusAgentRefundticketConfirmAPIRequest) GetParamAgentConfirmRefundRQ() *AgentConfirmRefundRq {
	return r._paramAgentConfirmRefundRQ
}

var poolTaobaoBusAgentRefundticketConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusAgentRefundticketConfirmRequest()
	},
}

// GetTaobaoBusAgentRefundticketConfirmRequest 从 sync.Pool 获取 TaobaoBusAgentRefundticketConfirmAPIRequest
func GetTaobaoBusAgentRefundticketConfirmAPIRequest() *TaobaoBusAgentRefundticketConfirmAPIRequest {
	return poolTaobaoBusAgentRefundticketConfirmAPIRequest.Get().(*TaobaoBusAgentRefundticketConfirmAPIRequest)
}

// ReleaseTaobaoBusAgentRefundticketConfirmAPIRequest 将 TaobaoBusAgentRefundticketConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusAgentRefundticketConfirmAPIRequest(v *TaobaoBusAgentRefundticketConfirmAPIRequest) {
	v.Reset()
	poolTaobaoBusAgentRefundticketConfirmAPIRequest.Put(v)
}
