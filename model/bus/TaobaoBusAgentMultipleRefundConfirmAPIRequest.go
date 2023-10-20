package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentMultipleRefundConfirmAPIRequest 综合交通多次退款接口 API请求
// taobao.bus.agent.multiple.refund.confirm
//
// 此接口支持多次按照单客进行多次退款操作，只进行退款操作。
type TaobaoBusAgentMultipleRefundConfirmAPIRequest struct {
	model.Params
	// 入参
	_paramAgentMultipleRefundRQ *AgentMultipleRefundRq
}

// NewTaobaoBusAgentMultipleRefundConfirmRequest 初始化TaobaoBusAgentMultipleRefundConfirmAPIRequest对象
func NewTaobaoBusAgentMultipleRefundConfirmRequest() *TaobaoBusAgentMultipleRefundConfirmAPIRequest {
	return &TaobaoBusAgentMultipleRefundConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusAgentMultipleRefundConfirmAPIRequest) Reset() {
	r._paramAgentMultipleRefundRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentMultipleRefundConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.multiple.refund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentMultipleRefundConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusAgentMultipleRefundConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentMultipleRefundRQ is ParamAgentMultipleRefundRQ Setter
// 入参
func (r *TaobaoBusAgentMultipleRefundConfirmAPIRequest) SetParamAgentMultipleRefundRQ(_paramAgentMultipleRefundRQ *AgentMultipleRefundRq) error {
	r._paramAgentMultipleRefundRQ = _paramAgentMultipleRefundRQ
	r.Set("param_agent_multiple_refund_r_q", _paramAgentMultipleRefundRQ)
	return nil
}

// GetParamAgentMultipleRefundRQ ParamAgentMultipleRefundRQ Getter
func (r TaobaoBusAgentMultipleRefundConfirmAPIRequest) GetParamAgentMultipleRefundRQ() *AgentMultipleRefundRq {
	return r._paramAgentMultipleRefundRQ
}

var poolTaobaoBusAgentMultipleRefundConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusAgentMultipleRefundConfirmRequest()
	},
}

// GetTaobaoBusAgentMultipleRefundConfirmRequest 从 sync.Pool 获取 TaobaoBusAgentMultipleRefundConfirmAPIRequest
func GetTaobaoBusAgentMultipleRefundConfirmAPIRequest() *TaobaoBusAgentMultipleRefundConfirmAPIRequest {
	return poolTaobaoBusAgentMultipleRefundConfirmAPIRequest.Get().(*TaobaoBusAgentMultipleRefundConfirmAPIRequest)
}

// ReleaseTaobaoBusAgentMultipleRefundConfirmAPIRequest 将 TaobaoBusAgentMultipleRefundConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusAgentMultipleRefundConfirmAPIRequest(v *TaobaoBusAgentMultipleRefundConfirmAPIRequest) {
	v.Reset()
	poolTaobaoBusAgentMultipleRefundConfirmAPIRequest.Put(v)
}
