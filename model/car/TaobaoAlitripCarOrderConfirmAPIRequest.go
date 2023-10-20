package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderConfirmAPIRequest 司机应答API API请求
// taobao.alitrip.car.order.confirm
//
// 航旅事业群-度假事业部-旅行用车项目组对外部服务商提供的司机应答回调接口
type TaobaoAlitripCarOrderConfirmAPIRequest struct {
	model.Params
	// 请求对象
	_paramOrderConfirm *OrderConfirm
}

// NewTaobaoAlitripCarOrderConfirmRequest 初始化TaobaoAlitripCarOrderConfirmAPIRequest对象
func NewTaobaoAlitripCarOrderConfirmRequest() *TaobaoAlitripCarOrderConfirmAPIRequest {
	return &TaobaoAlitripCarOrderConfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarOrderConfirmAPIRequest) Reset() {
	r._paramOrderConfirm = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarOrderConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderConfirm is ParamOrderConfirm Setter
// 请求对象
func (r *TaobaoAlitripCarOrderConfirmAPIRequest) SetParamOrderConfirm(_paramOrderConfirm *OrderConfirm) error {
	r._paramOrderConfirm = _paramOrderConfirm
	r.Set("param_order_confirm", _paramOrderConfirm)
	return nil
}

// GetParamOrderConfirm ParamOrderConfirm Getter
func (r TaobaoAlitripCarOrderConfirmAPIRequest) GetParamOrderConfirm() *OrderConfirm {
	return r._paramOrderConfirm
}

var poolTaobaoAlitripCarOrderConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarOrderConfirmRequest()
	},
}

// GetTaobaoAlitripCarOrderConfirmRequest 从 sync.Pool 获取 TaobaoAlitripCarOrderConfirmAPIRequest
func GetTaobaoAlitripCarOrderConfirmAPIRequest() *TaobaoAlitripCarOrderConfirmAPIRequest {
	return poolTaobaoAlitripCarOrderConfirmAPIRequest.Get().(*TaobaoAlitripCarOrderConfirmAPIRequest)
}

// ReleaseTaobaoAlitripCarOrderConfirmAPIRequest 将 TaobaoAlitripCarOrderConfirmAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarOrderConfirmAPIRequest(v *TaobaoAlitripCarOrderConfirmAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarOrderConfirmAPIRequest.Put(v)
}
