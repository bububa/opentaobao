package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderAcceptAPIRequest 确认接单 API请求
// taobao.alitrip.car.order.accept
//
// 用来接收服务商确认接单信息
type TaobaoAlitripCarOrderAcceptAPIRequest struct {
	model.Params
	// 确认订单请求
	_paramOrderAccept *OrderAccept
}

// NewTaobaoAlitripCarOrderAcceptRequest 初始化TaobaoAlitripCarOrderAcceptAPIRequest对象
func NewTaobaoAlitripCarOrderAcceptRequest() *TaobaoAlitripCarOrderAcceptAPIRequest {
	return &TaobaoAlitripCarOrderAcceptAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarOrderAcceptAPIRequest) Reset() {
	r._paramOrderAccept = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderAcceptAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.accept"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderAcceptAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarOrderAcceptAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderAccept is ParamOrderAccept Setter
// 确认订单请求
func (r *TaobaoAlitripCarOrderAcceptAPIRequest) SetParamOrderAccept(_paramOrderAccept *OrderAccept) error {
	r._paramOrderAccept = _paramOrderAccept
	r.Set("param_order_accept", _paramOrderAccept)
	return nil
}

// GetParamOrderAccept ParamOrderAccept Getter
func (r TaobaoAlitripCarOrderAcceptAPIRequest) GetParamOrderAccept() *OrderAccept {
	return r._paramOrderAccept
}

var poolTaobaoAlitripCarOrderAcceptAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarOrderAcceptRequest()
	},
}

// GetTaobaoAlitripCarOrderAcceptRequest 从 sync.Pool 获取 TaobaoAlitripCarOrderAcceptAPIRequest
func GetTaobaoAlitripCarOrderAcceptAPIRequest() *TaobaoAlitripCarOrderAcceptAPIRequest {
	return poolTaobaoAlitripCarOrderAcceptAPIRequest.Get().(*TaobaoAlitripCarOrderAcceptAPIRequest)
}

// ReleaseTaobaoAlitripCarOrderAcceptAPIRequest 将 TaobaoAlitripCarOrderAcceptAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarOrderAcceptAPIRequest(v *TaobaoAlitripCarOrderAcceptAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarOrderAcceptAPIRequest.Put(v)
}
