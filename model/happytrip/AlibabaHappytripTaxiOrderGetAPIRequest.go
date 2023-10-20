package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiordergetAPIRequest 订单详情 API请求
// alibaba.happytrip.taxi.order.get
//
// 获取订单状态及详情
type AlibabahappytriptaxiordergetAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
}

// NewAlibabahappytriptaxiordergetRequest 初始化AlibabahappytriptaxiordergetAPIRequest对象
func NewAlibabahappytriptaxiordergetRequest() *AlibabahappytriptaxiordergetAPIRequest {
	return &AlibabahappytriptaxiordergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiordergetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiordergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiordergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabahappytriptaxiordergetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiordergetAPIRequest) GetOrderId() string {
	return r._orderId
}
