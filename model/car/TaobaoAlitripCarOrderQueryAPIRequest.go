package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripCarOrderQueryAPIRequest 飞猪订单状态查询接口 API请求
// taobao.alitrip.car.order.query
//
// 提供给直连商家查询在飞猪平台上产生的订单
type TaobaoAlitripCarOrderQueryAPIRequest struct {
	model.Params
	// 飞猪平台订单id
	_orderId string
}

// NewTaobaoAlitripCarOrderQueryRequest 初始化TaobaoAlitripCarOrderQueryAPIRequest对象
func NewTaobaoAlitripCarOrderQueryRequest() *TaobaoAlitripCarOrderQueryAPIRequest {
	return &TaobaoAlitripCarOrderQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripCarOrderQueryAPIRequest) Reset() {
	r._orderId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripCarOrderQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪平台订单id
func (r *TaobaoAlitripCarOrderQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoAlitripCarOrderQueryAPIRequest) GetOrderId() string {
	return r._orderId
}

var poolTaobaoAlitripCarOrderQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripCarOrderQueryRequest()
	},
}

// GetTaobaoAlitripCarOrderQueryRequest 从 sync.Pool 获取 TaobaoAlitripCarOrderQueryAPIRequest
func GetTaobaoAlitripCarOrderQueryAPIRequest() *TaobaoAlitripCarOrderQueryAPIRequest {
	return poolTaobaoAlitripCarOrderQueryAPIRequest.Get().(*TaobaoAlitripCarOrderQueryAPIRequest)
}

// ReleaseTaobaoAlitripCarOrderQueryAPIRequest 将 TaobaoAlitripCarOrderQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripCarOrderQueryAPIRequest(v *TaobaoAlitripCarOrderQueryAPIRequest) {
	v.Reset()
	poolTaobaoAlitripCarOrderQueryAPIRequest.Put(v)
}
