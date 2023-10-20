package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentRefundConfirmAPIRequest 汽车票退票和退款二合一接口 API请求
// taobao.bus.agent.refund.confirm
//
// 1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
type TaobaoBusAgentRefundConfirmAPIRequest struct {
	model.Params
	// 入参
	_paramAgentConfirmReturnAndRefundRQ *AgentConfirmReturnAndRefundRq
}

// NewTaobaoBusAgentRefundConfirmRequest 初始化TaobaoBusAgentRefundConfirmAPIRequest对象
func NewTaobaoBusAgentRefundConfirmRequest() *TaobaoBusAgentRefundConfirmAPIRequest {
	return &TaobaoBusAgentRefundConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusAgentRefundConfirmAPIRequest) Reset() {
	r._paramAgentConfirmReturnAndRefundRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentRefundConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.refund.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentRefundConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusAgentRefundConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentConfirmReturnAndRefundRQ is ParamAgentConfirmReturnAndRefundRQ Setter
// 入参
func (r *TaobaoBusAgentRefundConfirmAPIRequest) SetParamAgentConfirmReturnAndRefundRQ(_paramAgentConfirmReturnAndRefundRQ *AgentConfirmReturnAndRefundRq) error {
	r._paramAgentConfirmReturnAndRefundRQ = _paramAgentConfirmReturnAndRefundRQ
	r.Set("param_agent_confirm_return_and_refund_r_q", _paramAgentConfirmReturnAndRefundRQ)
	return nil
}

// GetParamAgentConfirmReturnAndRefundRQ ParamAgentConfirmReturnAndRefundRQ Getter
func (r TaobaoBusAgentRefundConfirmAPIRequest) GetParamAgentConfirmReturnAndRefundRQ() *AgentConfirmReturnAndRefundRq {
	return r._paramAgentConfirmReturnAndRefundRQ
}

var poolTaobaoBusAgentRefundConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusAgentRefundConfirmRequest()
	},
}

// GetTaobaoBusAgentRefundConfirmRequest 从 sync.Pool 获取 TaobaoBusAgentRefundConfirmAPIRequest
func GetTaobaoBusAgentRefundConfirmAPIRequest() *TaobaoBusAgentRefundConfirmAPIRequest {
	return poolTaobaoBusAgentRefundConfirmAPIRequest.Get().(*TaobaoBusAgentRefundConfirmAPIRequest)
}

// ReleaseTaobaoBusAgentRefundConfirmAPIRequest 将 TaobaoBusAgentRefundConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusAgentRefundConfirmAPIRequest(v *TaobaoBusAgentRefundConfirmAPIRequest) {
	v.Reset()
	poolTaobaoBusAgentRefundConfirmAPIRequest.Put(v)
}
