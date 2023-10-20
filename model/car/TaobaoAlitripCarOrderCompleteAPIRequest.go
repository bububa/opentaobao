package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderCompleteAPIRequest 服务完成API API请求
// taobao.alitrip.car.order.complete
//
// 用来接收服务商订单流程完成信息
type TaobaoAlitripCarOrderCompleteAPIRequest struct {
	model.Params
	// 服务完成API
	_paramOrderComplete *OrderComplete
}

// NewTaobaoAlitripCarOrderCompleteRequest 初始化TaobaoAlitripCarOrderCompleteAPIRequest对象
func NewTaobaoAlitripCarOrderCompleteRequest() *TaobaoAlitripCarOrderCompleteAPIRequest {
	return &TaobaoAlitripCarOrderCompleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarOrderCompleteAPIRequest) Reset() {
	r._paramOrderComplete = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderCompleteAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderCompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarOrderCompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderComplete is ParamOrderComplete Setter
// 服务完成API
func (r *TaobaoAlitripCarOrderCompleteAPIRequest) SetParamOrderComplete(_paramOrderComplete *OrderComplete) error {
	r._paramOrderComplete = _paramOrderComplete
	r.Set("param_order_complete", _paramOrderComplete)
	return nil
}

// GetParamOrderComplete ParamOrderComplete Getter
func (r TaobaoAlitripCarOrderCompleteAPIRequest) GetParamOrderComplete() *OrderComplete {
	return r._paramOrderComplete
}

var poolTaobaoAlitripCarOrderCompleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarOrderCompleteRequest()
	},
}

// GetTaobaoAlitripCarOrderCompleteRequest 从 sync.Pool 获取 TaobaoAlitripCarOrderCompleteAPIRequest
func GetTaobaoAlitripCarOrderCompleteAPIRequest() *TaobaoAlitripCarOrderCompleteAPIRequest {
	return poolTaobaoAlitripCarOrderCompleteAPIRequest.Get().(*TaobaoAlitripCarOrderCompleteAPIRequest)
}

// ReleaseTaobaoAlitripCarOrderCompleteAPIRequest 将 TaobaoAlitripCarOrderCompleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarOrderCompleteAPIRequest(v *TaobaoAlitripCarOrderCompleteAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarOrderCompleteAPIRequest.Put(v)
}
