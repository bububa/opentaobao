package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripcarorderqueryAPIRequest 飞猪订单状态查询接口 API请求
// taobao.alitrip.car.order.query
//
// 提供给直连商家查询在飞猪平台上产生的订单
type TaobaoalitripcarorderqueryAPIRequest struct {
	model.Params
	// 飞猪平台订单id
	_orderId string
}

// NewTaobaoalitripcarorderqueryRequest 初始化TaobaoalitripcarorderqueryAPIRequest对象
func NewTaobaoalitripcarorderqueryRequest() *TaobaoalitripcarorderqueryAPIRequest {
	return &TaobaoalitripcarorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripcarorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripcarorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripcarorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪平台订单id
func (r *TaobaoalitripcarorderqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoalitripcarorderqueryAPIRequest) GetOrderId() string {
	return r._orderId
}
