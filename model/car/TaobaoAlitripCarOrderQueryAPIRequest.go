package car

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripCarOrderQueryAPIRequest
飞猪订单状态查询接口 API请求
taobao.alitrip.car.order.query

提供给直连商家查询在飞猪平台上产生的订单 */
type TaobaoAlitripCarOrderQueryAPIRequest struct {
	model.Params
	// 飞猪平台订单id
	_orderId string
}

// NewTaobaoAlitripCarOrderQueryRequest 初始化TaobaoAlitripCarOrderQueryAPIRequest对象
func NewTaobaoAlitripCarOrderQueryRequest() *TaobaoAlitripCarOrderQueryAPIRequest {
	return &TaobaoAlitripCarOrderQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripCarOrderQueryAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.car.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripCarOrderQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 飞猪平台订单id
func (r *TaobaoAlitripCarOrderQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r TaobaoAlitripCarOrderQueryAPIRequest) GetOrderId() string {
	return r._orderId
}
