package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarRentOrderCancelAPIRequest 租车-取消订单 API请求
// taobao.alitrip.car.rent.order.cancel
//
// 服务商主动取消用户订单或者拒绝取消订单.
type TaobaoAlitripCarRentOrderCancelAPIRequest struct {
	model.Params
	// 取消请求对象
	_param0 *RentProviderCancelRequest
}

// NewTaobaoAlitripCarRentOrderCancelRequest 初始化TaobaoAlitripCarRentOrderCancelAPIRequest对象
func NewTaobaoAlitripCarRentOrderCancelRequest() *TaobaoAlitripCarRentOrderCancelAPIRequest {
	return &TaobaoAlitripCarRentOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarRentOrderCancelAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarRentOrderCancelAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.rent.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarRentOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarRentOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 取消请求对象
func (r *TaobaoAlitripCarRentOrderCancelAPIRequest) SetParam0(_param0 *RentProviderCancelRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r TaobaoAlitripCarRentOrderCancelAPIRequest) GetParam0() *RentProviderCancelRequest {
	return r._param0
}

var poolTaobaoAlitripCarRentOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarRentOrderCancelRequest()
	},
}

// GetTaobaoAlitripCarRentOrderCancelRequest 从 sync.Pool 获取 TaobaoAlitripCarRentOrderCancelAPIRequest
func GetTaobaoAlitripCarRentOrderCancelAPIRequest() *TaobaoAlitripCarRentOrderCancelAPIRequest {
	return poolTaobaoAlitripCarRentOrderCancelAPIRequest.Get().(*TaobaoAlitripCarRentOrderCancelAPIRequest)
}

// ReleaseTaobaoAlitripCarRentOrderCancelAPIRequest 将 TaobaoAlitripCarRentOrderCancelAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarRentOrderCancelAPIRequest(v *TaobaoAlitripCarRentOrderCancelAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarRentOrderCancelAPIRequest.Put(v)
}
