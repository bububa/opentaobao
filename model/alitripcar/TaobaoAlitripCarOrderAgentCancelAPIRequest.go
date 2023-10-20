package alitripcar

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderAgentCancelAPIRequest 司机或客服取消订单 API请求
// taobao.alitrip.car.order.agent.cancel
//
// 司机或客服取消订单后通知飞猪订单取消
type TaobaoAlitripCarOrderAgentCancelAPIRequest struct {
	model.Params
	// 取消对象
	_paramOrderCancel *OrderCancel
}

// NewTaobaoAlitripCarOrderAgentCancelRequest 初始化TaobaoAlitripCarOrderAgentCancelAPIRequest对象
func NewTaobaoAlitripCarOrderAgentCancelRequest() *TaobaoAlitripCarOrderAgentCancelAPIRequest {
	return &TaobaoAlitripCarOrderAgentCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarOrderAgentCancelAPIRequest) Reset() {
	r._paramOrderCancel = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderAgentCancelAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.agent.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderAgentCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarOrderAgentCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderCancel is ParamOrderCancel Setter
// 取消对象
func (r *TaobaoAlitripCarOrderAgentCancelAPIRequest) SetParamOrderCancel(_paramOrderCancel *OrderCancel) error {
	r._paramOrderCancel = _paramOrderCancel
	r.Set("param_order_cancel", _paramOrderCancel)
	return nil
}

// GetParamOrderCancel ParamOrderCancel Getter
func (r TaobaoAlitripCarOrderAgentCancelAPIRequest) GetParamOrderCancel() *OrderCancel {
	return r._paramOrderCancel
}

var poolTaobaoAlitripCarOrderAgentCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarOrderAgentCancelRequest()
	},
}

// GetTaobaoAlitripCarOrderAgentCancelRequest 从 sync.Pool 获取 TaobaoAlitripCarOrderAgentCancelAPIRequest
func GetTaobaoAlitripCarOrderAgentCancelAPIRequest() *TaobaoAlitripCarOrderAgentCancelAPIRequest {
	return poolTaobaoAlitripCarOrderAgentCancelAPIRequest.Get().(*TaobaoAlitripCarOrderAgentCancelAPIRequest)
}

// ReleaseTaobaoAlitripCarOrderAgentCancelAPIRequest 将 TaobaoAlitripCarOrderAgentCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarOrderAgentCancelAPIRequest(v *TaobaoAlitripCarOrderAgentCancelAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarOrderAgentCancelAPIRequest.Put(v)
}
