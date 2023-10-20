package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniordergoodsreadyAPIRequest 备货完成 API请求
// taobao.omni.order.goods.ready
//
// 备货完成
type TaobaoomniordergoodsreadyAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewTaobaoomniordergoodsreadyRequest 初始化TaobaoomniordergoodsreadyAPIRequest对象
func NewTaobaoomniordergoodsreadyRequest() *TaobaoomniordergoodsreadyAPIRequest {
	return &TaobaoomniordergoodsreadyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniordergoodsreadyAPIRequest) GetApiMethodName() string {
	return "taobao.omni.order.goods.ready"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniordergoodsreadyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniordergoodsreadyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoomniordergoodsreadyAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoomniordergoodsreadyAPIRequest) GetOrderId() int64 {
	return r._orderId
}
