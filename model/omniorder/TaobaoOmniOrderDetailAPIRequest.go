package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniOrderDetailAPIRequest 全渠道订单详情 API请求
// taobao.omni.order.detail
//
// 全渠道订单详情
type TaobaoOmniOrderDetailAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewTaobaoOmniOrderDetailRequest 初始化TaobaoOmniOrderDetailAPIRequest对象
func NewTaobaoOmniOrderDetailRequest() *TaobaoOmniOrderDetailAPIRequest {
	return &TaobaoOmniOrderDetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniOrderDetailAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniOrderDetailAPIRequest) GetApiMethodName() string {
	return "taobao.omni.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniOrderDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniOrderDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoOmniOrderDetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoOmniOrderDetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoOmniOrderDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniOrderDetailRequest()
	},
}

// GetTaobaoOmniOrderDetailRequest 从 sync.Pool 获取 TaobaoOmniOrderDetailAPIRequest
func GetTaobaoOmniOrderDetailAPIRequest() *TaobaoOmniOrderDetailAPIRequest {
	return poolTaobaoOmniOrderDetailAPIRequest.Get().(*TaobaoOmniOrderDetailAPIRequest)
}

// ReleaseTaobaoOmniOrderDetailAPIRequest 将 TaobaoOmniOrderDetailAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniOrderDetailAPIRequest(v *TaobaoOmniOrderDetailAPIRequest) {
	v.Reset()
	poolTaobaoOmniOrderDetailAPIRequest.Put(v)
}
