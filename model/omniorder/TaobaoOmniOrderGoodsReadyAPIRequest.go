package omniorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniOrderGoodsReadyAPIRequest 备货完成 API请求
// taobao.omni.order.goods.ready
//
// 备货完成
type TaobaoOmniOrderGoodsReadyAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewTaobaoOmniOrderGoodsReadyRequest 初始化TaobaoOmniOrderGoodsReadyAPIRequest对象
func NewTaobaoOmniOrderGoodsReadyRequest() *TaobaoOmniOrderGoodsReadyAPIRequest {
	return &TaobaoOmniOrderGoodsReadyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOmniOrderGoodsReadyAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniOrderGoodsReadyAPIRequest) GetApiMethodName() string {
	return "taobao.omni.order.goods.ready"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniOrderGoodsReadyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOmniOrderGoodsReadyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoOmniOrderGoodsReadyAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoOmniOrderGoodsReadyAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolTaobaoOmniOrderGoodsReadyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOmniOrderGoodsReadyRequest()
	},
}

// GetTaobaoOmniOrderGoodsReadyRequest 从 sync.Pool 获取 TaobaoOmniOrderGoodsReadyAPIRequest
func GetTaobaoOmniOrderGoodsReadyAPIRequest() *TaobaoOmniOrderGoodsReadyAPIRequest {
	return poolTaobaoOmniOrderGoodsReadyAPIRequest.Get().(*TaobaoOmniOrderGoodsReadyAPIRequest)
}

// ReleaseTaobaoOmniOrderGoodsReadyAPIRequest 将 TaobaoOmniOrderGoodsReadyAPIRequest 放入 sync.Pool
func ReleaseTaobaoOmniOrderGoodsReadyAPIRequest(v *TaobaoOmniOrderGoodsReadyAPIRequest) {
	v.Reset()
	poolTaobaoOmniOrderGoodsReadyAPIRequest.Put(v)
}
