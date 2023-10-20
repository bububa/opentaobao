package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradedrugordergetAPIRequest 查看订单详情 API请求
// taobao.trade.drug.order.get
//
// 商家查看订单详情
type TaobaotradedrugordergetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewTaobaotradedrugordergetRequest 初始化TaobaotradedrugordergetAPIRequest对象
func NewTaobaotradedrugordergetRequest() *TaobaotradedrugordergetAPIRequest {
	return &TaobaotradedrugordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradedrugordergetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradedrugordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradedrugordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaotradedrugordergetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaotradedrugordergetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
