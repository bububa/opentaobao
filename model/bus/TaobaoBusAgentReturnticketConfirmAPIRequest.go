package bus

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentReturnticketConfirmAPIRequest 商家回调退票 API请求
// taobao.bus.agent.returnticket.confirm
//
// 商家通过TOP接口调用来回传退票状态
type TaobaoBusAgentReturnticketConfirmAPIRequest struct {
	model.Params
	// 退票入参
	_paramAgentConfirmReturnRQ *AgentConfirmReturnRq
}

// NewTaobaoBusAgentReturnticketConfirmRequest 初始化TaobaoBusAgentReturnticketConfirmAPIRequest对象
func NewTaobaoBusAgentReturnticketConfirmRequest() *TaobaoBusAgentReturnticketConfirmAPIRequest {
	return &TaobaoBusAgentReturnticketConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBusAgentReturnticketConfirmAPIRequest) Reset() {
	r._paramAgentConfirmReturnRQ = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentReturnticketConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.bus.agent.returnticket.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentReturnticketConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBusAgentReturnticketConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamAgentConfirmReturnRQ is ParamAgentConfirmReturnRQ Setter
// 退票入参
func (r *TaobaoBusAgentReturnticketConfirmAPIRequest) SetParamAgentConfirmReturnRQ(_paramAgentConfirmReturnRQ *AgentConfirmReturnRq) error {
	r._paramAgentConfirmReturnRQ = _paramAgentConfirmReturnRQ
	r.Set("param_agent_confirm_return_r_q", _paramAgentConfirmReturnRQ)
	return nil
}

// GetParamAgentConfirmReturnRQ ParamAgentConfirmReturnRQ Getter
func (r TaobaoBusAgentReturnticketConfirmAPIRequest) GetParamAgentConfirmReturnRQ() *AgentConfirmReturnRq {
	return r._paramAgentConfirmReturnRQ
}

var poolTaobaoBusAgentReturnticketConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBusAgentReturnticketConfirmRequest()
	},
}

// GetTaobaoBusAgentReturnticketConfirmRequest 从 sync.Pool 获取 TaobaoBusAgentReturnticketConfirmAPIRequest
func GetTaobaoBusAgentReturnticketConfirmAPIRequest() *TaobaoBusAgentReturnticketConfirmAPIRequest {
	return poolTaobaoBusAgentReturnticketConfirmAPIRequest.Get().(*TaobaoBusAgentReturnticketConfirmAPIRequest)
}

// ReleaseTaobaoBusAgentReturnticketConfirmAPIRequest 将 TaobaoBusAgentReturnticketConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoBusAgentReturnticketConfirmAPIRequest(v *TaobaoBusAgentReturnticketConfirmAPIRequest) {
	v.Reset()
	poolTaobaoBusAgentReturnticketConfirmAPIRequest.Put(v)
}
